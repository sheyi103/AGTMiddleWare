package util

import (
	"time"

	"github.com/spf13/viper"
)

//config stores all configuaration of the application.
// the values are read by viper from a config file or environment variable
type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACESS_TOKEN_DURATION"`
	MADAPI_CLIENT_ID     string        `mapstructure:"MADAPI_CLIENT_ID"`
	MADAPI_CLIENT_SECRET string        `mapstructure:"MADAPI_CLIENT_SECRET"`
	AGT_SMS_NOTIFY_URL   string        `mapstructure:"AGT_SMS_NOTIFY_URL"`
	AGT_USSD_NOTIFY_URL  string        `mapstructure:"AGT_USSD_NOTIFY_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
