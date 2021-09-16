package cli

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "pinger [command]",
		Short: "Pinger application",
		Long:  `Application for pinger another host`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	if len(args) == 0 {
		// 		cmd.Help()
		// 		os.Exit(0)
		// 	}
		// },
	}
)

func GetRootCmd() *cobra.Command {
	return rootCmd
}
