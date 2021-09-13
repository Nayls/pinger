package config

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
