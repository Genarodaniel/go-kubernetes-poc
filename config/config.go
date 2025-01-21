package config

import "github.com/spf13/viper"

type Env struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	GinMode    string `mapstructure:"GIN_MODE"`
	VIACEPURL  string `mapstructure:"VIACEP_URL"`
}

var Config Env

func Load() error {
	viper.AddConfigPath("../")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Config); err != nil {
		return err
	}

	return nil
}
