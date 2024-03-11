package jsonschema_test

import (
	"encoding/json"
	"fmt"
	jsonschema "jsonschema/drafts/4"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type testCase struct {
	Schema jsonschema.Schema
	Input  string
	Errors []string
}

func TestString(t *testing.T) {
	cases := []testCase{}

	err := filepath.Walk("./testcases/string", func(path string, info os.FileInfo, err error) error {
		if err == nil && info.IsDir() {
			ns := jsonschema.New()
			schema, err := ns.Read(filepath.Join(path, "schema.json"))

			if err != nil {
				return nil
			}

			input, err := os.ReadFile(filepath.Join(path, "input.txt"))

			if err != nil {
				return nil
			}

			text, err := os.ReadFile(filepath.Join(path, "errors.json"))

			if err != nil {
				if err != os.ErrNotExist {
					return nil
				}
			}

			errors := []string{}

			if text != nil {
				err := json.Unmarshal(text, &errors)

				if err != nil {
					return err
				}
			}

			cases = append(cases, testCase{
				Schema: schema,
				Input:  strings.TrimSpace(string(input)),
				Errors: errors,
			})
		}

		return err
	})

	if err != nil {
		t.Error(err)
	}

	for _, testcase := range cases {
		t.Run(testcase.Schema.GetTitle(), func(t *testing.T) {
			ns := jsonschema.New().AddSchema(testcase.Schema)
			errs := ns.Compile(testcase.Schema.GetID())

			t.Log(testcase.Schema)
			t.Logf(`expected: "%v"`, testcase.Errors)

			if len(errs) > 0 {
				if fmt.Sprint(testcase.Errors) != fmt.Sprint(errs) {
					t.Errorf(`received: "%v"`, errs)
				}

				return
			}

			errs = ns.Validate(testcase.Schema.GetID(), testcase.Input)

			if fmt.Sprint(testcase.Errors) != fmt.Sprint(errs) {
				t.Errorf(`received: "%v"`, errs)
			}
		})
	}
}
