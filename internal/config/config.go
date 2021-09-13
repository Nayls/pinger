package config

import (
	"github.com/spf13/cobra"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port" validate:"required"`
	} `yaml:"server"`

	Logger struct {
		Level string `yaml:"level" validate:"required"`
	} `yaml:"logger"`

	Database struct {
		Host string `yaml:"host" validate:"required"`
		Port string `yaml:"port" validate:"required"`
	} `yaml:"database"`
}

var (
	rootCmd = &cobra.Command{
		Use:   "pinger [command]",
		Short: "Pinger application",
		Long:  `Application for pinger another host`,
		Run: func(cmd *cobra.Command, args []string) {
			// if len(args) == 0 {
			// 	cmd.Help()
			// 	os.Exit(0)
			// }
			// do actual work
		},
	}
)

func GetRootCmd() *cobra.Command {
	return rootCmd
}
