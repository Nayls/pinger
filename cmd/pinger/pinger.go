package pinger

import (
	"log"
	"os"

	"github.com/Nayls/pinger/cmd/completion"
	"github.com/Nayls/pinger/cmd/generate"
	"github.com/Nayls/pinger/cmd/server"
	"github.com/Nayls/pinger/internal/cli"
	"github.com/spf13/cobra"
)

func InitCobraConfig() *cobra.Command {
	cobra.OnInitialize(func() {
		// if _, err := os.Stat("./docs"); os.IsNotExist(err) {
		// 	if err := os.Mkdir("./docs", 0755); err != nil {
		// 		log.Fatal(err)
		// 	}
		// }
		// if err := doc.GenMarkdownTree(rootCmd, "./docs"); err != nil {
		// 	log.Fatal(err)
		// }
	})

	rootCmd := cli.GetRootCmd()

	// Add generate command
	rootCmd.AddCommand(generate.GetGenerateCmd())

	// Add server command
	rootCmd.AddCommand(server.GetServerCmd())

	// Add completion command
	rootCmd.AddCommand(completion.GetCompletionCmd())

	return rootCmd
}

func init() {
	if err := InitCobraConfig().Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
