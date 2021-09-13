package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var config Config

func ViperServerConfiguration() {
	// Server default configurations
	viper.SetDefault("server.host", "")
	viper.SetDefault("server.port", "8080")

	// Logger default configurations
	viper.SetDefault("logger.level", "debug")

	// Database default configurations
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5439")

	// Read server-config.yaml
	if viper.IsSet("config") && viper.Get("config") == 0 {
		log.Print("config is set")
	} else {
		viper.SetConfigName("server-config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./configs")
		viper.AddConfigPath(".")
	}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("config file not found error - %v", err)
		} else if _, ok := err.(viper.ConfigMarshalError); ok {
			log.Printf("config marshal error - %v", err)
		} else {
			log.Printf("config filer read error - %v", err)
		}
	}

	// Enable watching config file
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Print("config file changed:", e.Name)
	})
	viper.WatchConfig()

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
	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("unmasrshal config error - %v", err)
	}

	if err := validator.New().Struct(&config); err != nil {
		log.Printf("eror in validation config - %v", err)
	}
}
