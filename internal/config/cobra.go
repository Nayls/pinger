package config

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "pinger",
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

func CobraServerConfiguration() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("config", "c", "", "config file path (default is ./server-config.yaml)")
	if err := viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config")); err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().String("server-host", "", "config file (server host \"\")")
	rootCmd.PersistentFlags().String("server-port", "", "config file (server port \"8080\")")

	if err := viper.BindPFlag("server.host", rootCmd.PersistentFlags().Lookup("server-host")); err != nil {
		log.Fatal(err)
	}
	if err := viper.BindPFlag("server.port", rootCmd.PersistentFlags().Lookup("server-port")); err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().String("logger-level", "", "logger level (default is \"debug\")")

	if err := viper.BindPFlag("logger.level", rootCmd.PersistentFlags().Lookup("logger-level")); err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().String("database-host", "", "database host (default is \"localhost\")")
	rootCmd.PersistentFlags().String("database-port", "", "database port (default is \"5439\")")
	if err := viper.BindPFlag("database.host", rootCmd.PersistentFlags().Lookup("database-host")); err != nil {
		log.Fatal(err)
	}
	if err := viper.BindPFlag("database.port", rootCmd.PersistentFlags().Lookup("database-port")); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	cfgFile, _ := rootCmd.Flags().GetString("config")
	if cfgFile != "" {
		log.Print("Config:", cfgFile)
	}
}
