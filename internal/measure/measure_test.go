package measure_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/cmontemuino/tokeval/internal/measure"
)

func TestNewCmdMeasure(t *testing.T) {
	// Mock the runner function
	mockRunner := func(args []string, w io.Writer) error {
		_, err := w.Write([]byte("hello world"))
		return err
	}

	cmd := measure.NewCmdMeasure("tokeval", measure.WithRunner(mockRunner))

	var outBuf, errBuf bytes.Buffer
	cmd.SetOut(&outBuf)
	cmd.SetErr(&errBuf)

	cmd.SetArgs([]string{"--tokenizer", "custom_test", "--", "echo", "hello world"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Command execution failed %v", err)
	}

	// runMeasure explicitly prints the summary info to Stderr.
	gotStderr := errBuf.String()

	if !strings.Contains(gotStderr, "custom_test") {
		t.Errorf("Expected output to contain custom_test, got:\n%s", gotStderr)
	}

	if !strings.Contains(gotStderr, "Bytes:   11") {
		t.Errorf("Expected output to contain Bytes: 11, got:\n%s", gotStderr)
	}
}
