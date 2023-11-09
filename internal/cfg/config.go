package cfg

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string
	Host          string
	Port          string
	LinkLength    int
	Username      string
	Password      string
	DbName        string
	SSLMode       string
}

func Init(path string) {
	if err := initConfig(path); err != nil {
		log.Fatalf("faliled to initialize configs: %s", err)
	}
}

// TODO:При тестах не может найти файл,если не указать абсолютный путь
func initConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}

func ServConfig() Config {
	servCfg := Config{
		ServerAddress: viper.GetString("server.host") + ":" + viper.GetString("server.port"),
		Port:          viper.GetString("server.port"),
		LinkLength:    viper.GetInt("server.linkLength"),
	}
	return servCfg
}

func DbConfig() Config {
	dbCfg := Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DbName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
	return dbCfg
}
