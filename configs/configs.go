package configs

import (
	"fast-search/pkg/utils"
	"os"

	"github.com/spf13/viper"
)

type Configs struct {
	GRPCConection string `mapstructure:"grpc_conection" validate:"required"`
	environment   string
}

func loadConfig(path string) Configs {
	env := os.Getenv("ENVIRONMENT")
	if env == "PRODUCTION" {
		viper.SetConfigName("env.production")
	} else if env == "STAGING" {
		viper.SetConfigName("env.staging")
	} else {
		env = "LOCAL"
		viper.SetConfigName("env.local")
	}
	viper.SetConfigType("toml")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var config Configs
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	config.environment = env
	return config
}

func (c *Configs) GetEnvironment() string {
	return c.environment
}

func (c *Configs) Validate() error {
	return utils.ValidateStruct(c)
}

func NewConfigs(workDir string) *Configs {
	config := loadConfig(workDir)
	err := config.Validate()
	if err != nil {
		panic(err)
	}
	return &config
}
