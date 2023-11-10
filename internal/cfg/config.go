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
	Format        bool
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

func ServConfig(s string) Config {
	if s == "gRPC" {
		servCfg := Config{
			ServerAddress: viper.GetString("gRPCserver.host") + ":" + viper.GetString("gRPCserver.port"),
			Port:          viper.GetString("gRPCserver.port"),
			LinkLength:    viper.GetInt("gRPCserver.linkLength"),
		}
		return servCfg
	} else {
		servCfg := Config{
			ServerAddress: viper.GetString("HTTPserver.host") + ":" + viper.GetString("HTTPserver.port"),
			Port:          viper.GetString("HTTPserver.port"),
			LinkLength:    viper.GetInt("HTTPserver.linkLength"),
		}
		return servCfg
	}
}

func DbConfig() Config {
	dbCfg := Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DbName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Format:   viper.GetBool("db.format"),
	}
	return dbCfg
}
