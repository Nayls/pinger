package pinger

import (
	"github.com/Nayls/pinger/cmd/completion"
	"github.com/Nayls/pinger/cmd/generate"
	"github.com/Nayls/pinger/cmd/server"
	"github.com/Nayls/pinger/internal/cli"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "pinger [command]",
		Short:   "Pinger application",
		Long:    `Application for pinger another host`,
		Version: "v0.2",
		// Run: func(cmd *cobra.Command, args []string) {
		// 	if len(args) == 0 {
		// 		cmd.Help()
		// 		os.Exit(0)
		// 	}
		// },
	}
)

func GetPingerCmd() *cobra.Command {
	return rootCmd
}

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
