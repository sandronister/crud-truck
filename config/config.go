package config

import (
	"github.com/spf13/viper"
)

type Conf struct {
	WebPort    string   `mapstructure:"WEB_PORT"`
	DBDriver   string   `mapstructure:"DB_DRIVER"`
	DBHost     string   `mapstructure:"DB_HOST"`
	DBPort     string   `mapstructure:"DB_PORT"`
	DBUser     string   `mapstructure:"DB_USER"`
	DBPassword string   `mapstructure:"DB_PASSWORD"`
	DBName     string   `mapstructure:"DB_NAME"`
	Ips        []string `mapstructure:"IPS"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil

}
