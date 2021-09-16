package generate

import (
	"log"
	"os"

	"github.com/Nayls/pinger/internal/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate [command]",
	Short: "Command generate",
	Long:  `Command for generate docs and etc`,
	// Run:          func(cmd *cobra.Command, args []string) {},
}

var generateCliDocCmd = &cobra.Command{
	Use:   "cli",
	Short: "Generate cli documentation",
	Long:  `Generate cli documentations`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat("./docs/cli"); os.IsNotExist(err) {
			if err := os.MkdirAll("./docs/cli", 0755); err != nil {
				log.Fatal(err)
			}
		}

		if err := doc.GenMarkdownTree(cli.GetRootCmd(), "./docs/cli"); err != nil {
			log.Fatal(err)
		}
	},
}

func GetGenerateCmd() *cobra.Command {
	return generateCmd
}

func init() {
	cobra.OnInitialize(initConfig)

	// Add subcommands
	generateCmd.AddCommand(generateCliDocCmd)
}

func initConfig() {}
