// Package measure includes the logic for measuring tokens.
package measure

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"

	"github.com/cmontemuino/tokeval/internal/runner"
	"github.com/spf13/cobra"
)

// commandRunner is used to inject a Runner so that NewCmdMeasure can be tested.
type commandRunner func(args []string, w io.Writer) error

// Option defines a functional configuration option for the measure command.
type Option func(*measureOptions)

// WithRunner is a functional option for testing purposes.
// It overrides the command runner.
func WithRunner(r func([]string, io.Writer) error) Option {
	return func(o *measureOptions) { o.runner = r }
}

// WithLevel is a functional option for setting the level var.
func WithLogLevel(level *slog.LevelVar) Option {
	return func(o *measureOptions) { o.logLevel = level }
}

// measureOptions holds configuration parse from flags, allowing to select different tokenizers.
type measureOptions struct {
	tokenizer string
	verbose   bool
	runner    commandRunner // Allows mocking in tests
	logLevel  *slog.LevelVar
}

// NewCmdMeasure creates the `measure` subcommand.
// It accepts optional configurations.
func NewCmdMeasure(parentCommand string, options ...Option) *cobra.Command {
	opts := &measureOptions{
		runner: runner.Run,
	}

	// Apply any options provided
	for _, opt := range options {
		opt(opts)
	}

	cmd := &cobra.Command{
		Use:   "measure -- <command>",
		Short: "Measure tokens used by a command.",
		Long: `Execute a command and measure the number of tokens in its output.
The command to be measured must be preceded by '--'.`,
		Example: `  # Measure tokens for a karmadactl command
  tokeval measure -- karmadactl get clusters

  # Measure tokens for a local script
  tokeval measure -- ./my-script.sh`,

		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runMeasure(cmd, opts, args)
		},
	}

	// Bind flags to the measureOptions
	cmd.Flags().StringVarP(&opts.tokenizer, "tokenizer", "t", "cl100k_base", "Tokenizer to use for counting")
	cmd.Flags().BoolVarP(&opts.verbose, "verbose", "v", false, "Enable verbose debug logs")

	return cmd
}

func runMeasure(cmd *cobra.Command, opts *measureOptions, args []string) error {
	if opts.verbose && opts.logLevel != nil {
		opts.logLevel.Set(slog.LevelDebug)
	}

	var buf bytes.Buffer

	// Duplicate output to both terminal and buffer
	if err := opts.runner(args, &buf); err != nil {
		return fmt.Errorf("command execution failed: %w", err)
	}

	// Count tokens based on bytes for now. Using tokenizers comes later.
	content := buf.Bytes()

	// Keep stdout clean. Print summary to stderr instead.
	cmd.PrintErrf("\n--- Token Measurement (%s) ---\n", opts.tokenizer)
	cmd.PrintErrf("Bytes:   %d\n", len(content))
	cmd.PrintErrf("Tokens: %d (mocked)\n", len(content)/4) // Mock count for now

	return nil
}
