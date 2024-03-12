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
		t.Run(testcase.Schema.GetTitle(), func(t *testing.T) {
			ns := ns.AddSchema(testcase.Schema.(T))
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

func readAll[T core.Schema](path string, ns core.Namespace[T]) ([]Case, error) {
	cases := []Case{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil && info.IsDir() {
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

			cases = append(cases, Case{
				Schema: schema,
				Input:  strings.TrimSpace(string(input)),
				Errors: errors,
			})
		}

		return err
	})

	return cases, err
}
