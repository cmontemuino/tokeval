// Package main is the entrypoint for tokeval.
package main

import (
	"log/slog"
	"os"

	"github.com/cmontemuino/tokeval/internal/cli"
)

func main() {
	// Using TextHandler in favor of JsonHandler because if a user redirect the output to a file
	// (e.g., `tokeval measure -- git log > log.txt), the JSON log will be written along with the
	// command output. Also use os.Stderr, leaving os.Stdout untouched so that it is safe to pipe
	// the command output to other tools.
	var logLevel slog.LevelVar
	logLevel.Set(slog.LevelWarn)
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: &logLevel,
	}))
	slog.SetDefault(logger)

	cmd := cli.NewCmdTokeval("tokeval", "tokeval", &logLevel)
	if err := cmd.Execute(); err != nil {
		slog.Error("execution failed", "err", err)
		os.Exit(1)
	}
}
