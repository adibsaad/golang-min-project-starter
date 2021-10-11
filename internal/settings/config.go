package settings

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const (
	EnvPrefix = "MY_APP_"
)

type DatabaseConfig struct {
	Url      string `mapstructure:"url"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Config struct {
	// Hold a ref to the viper object
	Viper *viper.Viper

	Database DatabaseConfig `mapstructure:"db"`
}

func NewConfig(cfgFile string) *Config {
	v := viper.New()
	genDefault(v)
	config := &Config{Viper: v}

	config.Viper.SetEnvPrefix(EnvPrefix)
	config.Viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "__")
	config.Viper.SetEnvKeyReplacer(replacer)
	config.Viper.SetTypeByDefaultValue(true)

	if cfgFile != "" {
		// Use config file from the flag.
		config.Viper.SetConfigFile(cfgFile)

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	err := config.Viper.Unmarshal(&config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return config
}

func genDefault(viper *viper.Viper) {
	viper.SetDefault("db.url", "http://127.0.0.1")
	viper.SetDefault("db.port", 3306)
	viper.SetDefault("db.database", "my-app-db")
	viper.SetDefault("db.username", "root")
	viper.SetDefault("db.password", "")
}
