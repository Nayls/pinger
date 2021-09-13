package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/nayls.cloud/ping.nayls.cloud/pinger/cmd/completion"
	"gitlab.com/nayls.cloud/ping.nayls.cloud/pinger/cmd/generate"
	"gitlab.com/nayls.cloud/ping.nayls.cloud/pinger/cmd/server"
	"gitlab.com/nayls.cloud/ping.nayls.cloud/pinger/internal/config"
)

func initCobraConfig() {
	rootCmd := config.GetRootCmd()

	// Add generate command
	rootCmd.AddCommand(generate.GetGenerateCmd())

	// Add server command
	rootCmd.AddCommand(server.GetServerCmd())

	// Add completion command
	rootCmd.AddCommand(completion.GetCompletionCmd())

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

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func init() {
	initCobraConfig()
}

func main() {}
