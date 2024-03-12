package test

import (
	"encoding/json"
	"fmt"
	"jsonschema/core"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func RunAll[T core.Schema](path string, ns core.Namespace[T], t *testing.T) {
	cases, err := readAll(path, ns)

	if err != nil {
		t.Error(err)
	}

	for _, testcase := range cases {
		id := testcase.Schema.GetID()

		if id == "" {
			t.Log(testcase.Schema)
			t.Error(`test schemas "id" is required`)
		}

		title := id

		if strings.HasPrefix(title, "/") {
			title = title[1:]
		}

		t.Run(title, func(t *testing.T) {
			ns := ns.AddSchema(testcase.Schema.(T))
			errs := ns.Compile(id)

			if len(errs) > 0 {
				if fmt.Sprint(testcase.Errors) != fmt.Sprint(errs) {
					t.Log(testcase.Schema)
					t.Logf(`expected: "%v"`, testcase.Errors)
					t.Errorf(`received: "%v"`, errs)
				}

				return
			}

			errs = ns.Validate(id, testcase.Input)

			if fmt.Sprint(testcase.Errors) != fmt.Sprint(errs) {
				t.Log(testcase.Schema)
				t.Logf(`expected: "%v"`, testcase.Errors)
				t.Errorf(`received: "%v"`, errs)
			}
		})
	}
}

func readAll[T core.Schema](path string, ns core.Namespace[T]) ([]Case, error) {
	cases := []Case{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil && info.IsDir() {
			schema, err := ns.Read(filepath.Join(path, "schema.json"))

			if err != nil {
				if os.IsNotExist(err) {
					return nil
				}

				return err
			}

			var input any
			input, err = os.ReadFile(filepath.Join(path, "input.json"))

			if err != nil {
				return err
			}

			err = json.Unmarshal(input.([]byte), &input)

			if err != nil {
				return err
			}

			text, err := os.ReadFile(filepath.Join(path, "errors.json"))

			if err != nil {
				if !os.IsNotExist(err) {
					return nil
				}
			}

			errors := []string{}

			if text != nil {
				err := json.Unmarshal(text, &errors)

				if err != nil {
					fmt.Println("d", path)
					return err
				}
			}

			cases = append(cases, Case{
				Schema: schema,
				Input:  input,
				Errors: errors,
			})
		}

		return err
	})

	return cases, err
}
