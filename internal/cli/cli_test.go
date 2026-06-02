package cli_test

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/cmontemuino/tokeval/internal/cli"
)

func TestNewCmdTokeval(t *testing.T) {
	var logLevel slog.LevelVar
	cmd := cli.NewCmdTokeval("tokeval", "tokeval", &logLevel)

	var outBuf, errBuf bytes.Buffer
	cmd.SetOut(&outBuf)
	cmd.SetErr(&errBuf)

	cmd.SetArgs([]string{})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Command execution failed: %v", err)
	}

	// Runnnig the root command with no arguments makes Cobra show the Help/Usage text. Cobra writes it to Stdout.
	gotOutput := outBuf.String()

	if !strings.Contains(gotOutput, "tokeval measures token counts.") {
		t.Errorf("Expected help outout, got:\n%s", gotOutput)
	}

	if !strings.Contains(gotOutput, "measure") {
		t.Errorf("Expected registered 'measure' subcommand not found, got:\n%s", gotOutput)
	}
}
