// Package cli assembles all commands offered by `tokeval`
package cli

import (
	"fmt"
	"log/slog"

	"github.com/cmontemuino/tokeval/internal/measure"
	"github.com/spf13/cobra"
)

const (
	rootCmdShort = "%s measures token counts."
	rootCmdLong  = "%s measures token counts."
)

// NewCmdTokeval creates the `tokeval` command.
func NewCmdTokeval(cmdUse, parentCommand string, logLevel *slog.LevelVar) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   cmdUse,
		Short: fmt.Sprintf(rootCmdShort, parentCommand),
		Long:  fmt.Sprintf(rootCmdLong, parentCommand),
		// Avoid printing the entire help in case of execution errors
		SilenceErrors: true,
		SilenceUsage:  true,

		RunE: runHelp,
	}

	rootCmd.AddCommand(measure.NewCmdMeasure(cmdUse, measure.WithLogLevel(logLevel)))

	return rootCmd
}

func runHelp(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}
