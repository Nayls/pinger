package server

import (
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/nayls.cloud/ping.nayls.cloud/pinger/internal/config"
)

// serverCmd represents the server command
var (
	serverCfgFile string
	serverCfg     config.Config

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Command for server managmenent",
		Long:  `Command for server managmenent and configurations`,
		Run:   func(cmd *cobra.Command, args []string) {},
	}
)

func GetServerCmd() *cobra.Command {
	return serverCmd
}

func init() {
	cobra.OnInitialize(initConfig)

	// Add flag for config file
	serverCmd.Flags().StringVarP(&serverCfgFile, "config", "c", "", "config file path (default is \"./config.yaml)\"")
	if err := viper.BindPFlag("server.config", serverCmd.Flags().Lookup("config")); err != nil {
		log.Fatal(err)
	}

	// Add flag for server
	serverCmd.Flags().StringVar(&serverCfg.Server.Host, "server-host", "", "server host (default is \"\")")
	serverCmd.Flags().StringVar(&serverCfg.Server.Port, "server-port", "", "server port (default is \"8080\")")
	if err := viper.BindPFlag("server.host", serverCmd.Flags().Lookup("server-host")); err != nil {
		log.Fatal(err)
	}
	if err := viper.BindPFlag("server.port", serverCmd.Flags().Lookup("server-port")); err != nil {
		log.Fatal(err)
	}

	// Add flag for logger
	serverCmd.Flags().StringVar(&serverCfg.Logger.Level, "logger-level", "", "logger level (default is \"debug\")")
	if err := viper.BindPFlag("logger.level", serverCmd.Flags().Lookup("logger-level")); err != nil {
		log.Fatal(err)
	}

	// Add flag for database
	serverCmd.Flags().StringVar(&serverCfg.Database.Host, "database-host", "", "database host (default is \"localhost\")")
	serverCmd.Flags().StringVar(&serverCfg.Database.Port, "database-port", "", "database port (default is \"5439\")")
	if err := viper.BindPFlag("database.host", serverCmd.Flags().Lookup("database-host")); err != nil {
		log.Fatal(err)
	}
	if err := viper.BindPFlag("database.port", serverCmd.Flags().Lookup("database-port")); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	// Server default configurations
	viper.SetDefault("server.host", "")
	viper.SetDefault("server.port", "8080")

	// Logger default configurations
	viper.SetDefault("logger.level", "debug")

	// Database default configurations
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5439")

	// Read server-config.yaml
	if viper.IsSet("server.config") && len(viper.GetString("server.config")) >= 0 {
		filename := filepath.Clean(viper.GetString("server.config"))
		ext := filepath.Ext(filename)[1:]
		name := filepath.Base(filename)[0 : len(filepath.Base(filename))-len(ext)-1]
		path := filepath.Clean(filename)[0 : len(filename)-len(name)-len(ext)-1]

		viper.SetConfigName(name)
		viper.SetConfigType(ext)
		viper.AddConfigPath(path)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./configs")
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err == nil {
		log.Print("Read server config - ", viper.ConfigFileUsed())
		// Enable watching config file
		viper.OnConfigChange(func(e fsnotify.Event) {
			log.Print("config file changed:", e.Name)
		})
		viper.WatchConfig()
	} else {
		log.Print("Server config is not read, uses default value")
	}

	// Read env in system environment
	viper.SetEnvPrefix("pinger")
	viper.AutomaticEnv()

	// Apply system environment variable in struct config
	// PINGER_SERVER_HOST -> server_host -> server.host
	if viper.IsSet("server_host") {
		viper.Set("server.host", viper.Get("server_host"))
	}
	if viper.IsSet("server_port") {
		viper.Set("server.port", viper.Get("server_port"))
	}
	if viper.IsSet("logger_level") {
		viper.Set("logger.level", viper.Get("logger_level"))
	}
	if viper.IsSet("database_host") {
		viper.Set("database.host", viper.Get("database_host"))
	}
	if viper.IsSet("database_port") {
		viper.Set("database.port", viper.Get("database_port"))
	}

	// Unmarshal config and validate
	var configuration config.Config
	if err := viper.Unmarshal(&configuration); err != nil {
		log.Printf("unmasrshal config error - %v", err)
	}

	if err := validator.New().Struct(&configuration); err != nil {
		log.Printf("eror in validation config - %v", err)
	}
}
