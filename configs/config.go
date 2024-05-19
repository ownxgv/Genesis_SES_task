package configspackage configs

import "github.com/spf13/viper"

type Config struct {
	Port        string `mapstructure:"PORT"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBUser      string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	DBName      string `mapstructure:"DB_NAME"`
	EmailSender string `mapstructure:"EMAIL_SENDER"`
	EmailPass   string `mapstructure:"EMAIL_PASS"`
}

var Config *Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	Config = &Config{}
	err = viper.Unmarshal(Config)
	if err != nil {
		panic(err)
	}
}
