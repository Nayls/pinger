package pinger

import (
	"github.com/Nayls/pinger/cmd/completion"
	"github.com/Nayls/pinger/cmd/generate"
	"github.com/Nayls/pinger/cmd/server"
	"github.com/Nayls/pinger/internal/cli"
	"github.com/spf13/cobra"
)

func InitCobraConfig() *cobra.Command {
	rootCmd := cli.GetRootCmd()

	// Add generate command
	rootCmd.AddCommand(generate.GetGenerateCmd())

	// Add server command
	rootCmd.AddCommand(server.GetServerCmd())

	// Add completion command
	rootCmd.AddCommand(completion.GetCompletionCmd())

	return rootCmd
}

func init() {}
