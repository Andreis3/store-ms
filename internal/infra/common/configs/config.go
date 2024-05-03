package configs

import (
	"github.com/spf13/viper"
)

type Conf struct {
	DBDriver        string `mapstructure:"DB_DRIVER"`
	DBHost          string `mapstructure:"DB_HOST"`
	DBPort          string `mapstructure:"DB_PORT"`
	DBUser          string `mapstructure:"DB_USER"`
	DBPassword      string `mapstructure:"DB_PASSWORD"`
	DBName          string `mapstructure:"DB_NAME"`
	ServerPort      string `mapstructure:"SERVER_PORT"`
	MaxConnections  string `mapstructure:"MAX_CONNECTIONS"`
	MinConnections  string `mapstructure:"MIN_CONNECTIONS"`
	MaxConnLifetime string `mapstructure:"MAX_CONN_LIFETIME"`
	MaxConnIdleTime string `mapstructure:"MAX_CONN_IDLE_TIME"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		cfg = &Conf{
			DBDriver:        viper.GetString("DB_DRIVER"),
			DBHost:          viper.GetString("DB_HOST"),
			DBPort:          viper.GetString("DB_PORT"),
			DBUser:          viper.GetString("DB_USER"),
			DBPassword:      viper.GetString("DB_PASSWORD"),
			DBName:          viper.GetString("DB_NAME"),
			ServerPort:      viper.GetString("SERVER_PORT"),
			MaxConnections:  viper.GetString("MAX_CONNECTIONS"),
			MinConnections:  viper.GetString("MIN_CONNECTIONS"),
			MaxConnLifetime: viper.GetString("MAX_CONN_LIFETIME"),
			MaxConnIdleTime: viper.GetString("MAX_CONN_IDLE_TIME"),
		}
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, err
}
