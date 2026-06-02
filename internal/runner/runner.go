// Package runner take a slice of strings captured by the `measure` subcommand and actual run it.
package runner

import (
	"io"
	"os"
	"os/exec"
)

// Run executes the external command specified by args.
// It duplicates the command's standard output to both stdout and the provided `tokenBuffer`.
// Standard input and standard error kept connected directl to parent OS streams.
func Run(args []string, tokenBuffer io.Writer) error {
	name := args[0]
	remainingArgs := args[1:]

	cmd := exec.Command(name, remainingArgs...)

	cmd.Stdout = io.MultiWriter(os.Stdout, tokenBuffer)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
