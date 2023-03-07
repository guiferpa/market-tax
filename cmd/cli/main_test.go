package main

import (
	"bytes"
	"io"
	"os/exec"
	"strings"
	"testing"
)

func TestIntegrationCLIByStdinSingleLine(t *testing.T) {
	sub := exec.Command("go", "run", "./main.go")

	stdin, err := sub.StdinPipe()
	if err != nil {
		t.Error(err)
		return
	}

	go func() {
		defer stdin.Close()

		value := `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000}, {"operation":"buy", "unit-cost":25.00, "quantity": 5000}, {"operation":"sell", "unit-cost":15.00, "quantity": 10000}]`
		if _, err := io.WriteString(stdin, value); err != nil {
			t.Error(err)
			return
		}
	}()

	output, err := sub.CombinedOutput()
	if err != nil {
		t.Error(err)
		return
	}

	expected := "[{\"tax\":0},{\"tax\":0},{\"tax\":0}]\n"
	buf := bytes.NewBuffer(output)
	if got, expected := buf.String(), expected; strings.Compare(got, expected) != 0 {
		t.Errorf("unexpected result for integration test, got: %s, expected: %s", got, expected)
		return
	}
}

func TestIntegrationCLIByStdinDoubleLine(t *testing.T) {
	sub := exec.Command("go", "run", "./main.go")

	stdin, err := sub.StdinPipe()
	if err != nil {
		t.Error(err)
		return
	}

	go func() {
		defer stdin.Close()

		value := `[{"operation":"buy", "unit-cost":10.00, "quantity": 100}, {"operation":"sell", "unit-cost":15.00, "quantity": 50}, {"operation":"sell", "unit-cost":15.00, "quantity": 50}]
[{"operation":"buy", "unit-cost":10.00, "quantity": 10000}, {"operation":"sell", "unit-cost":20.00, "quantity": 5000}, {"operation":"sell", "unit-cost":5.00, "quantity": 5000}]`
		if _, err := io.WriteString(stdin, value); err != nil {
			t.Error(err)
			return
		}
	}()

	output, err := sub.CombinedOutput()
	if err != nil {
		t.Error(err)
		return
	}

	expected := "[{\"tax\":0},{\"tax\":0},{\"tax\":0}]\n[{\"tax\":0},{\"tax\":10000},{\"tax\":0}]\n"
	buf := bytes.NewBuffer(output)
	if got, expected := buf.String(), expected; strings.Compare(got, expected) != 0 {
		t.Errorf("unexpected result for integration test, got: %s, expected: %s", got, expected)
		return
	}
}
