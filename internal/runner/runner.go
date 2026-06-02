// Package runner takes a slice of strings captured by the `measure` subcommand and runs it.
package runner

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// Run executes the external command specified by args.
// It duplicates the command's standard output to both stdout and the provided tokenBuffer.
// Standard input and standard error are connected directly to parent OS streams.
func Run(ctx context.Context, args []string, tokenBuffer io.Writer) error {
	if len(args) == 0 {
		return fmt.Errorf("runner: no command specified")
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = io.MultiWriter(os.Stdout, tokenBuffer)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
