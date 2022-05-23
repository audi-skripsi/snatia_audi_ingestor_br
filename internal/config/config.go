package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName    string
	AppAddress string

	Environment   Environment
	MongoDBConfig MongoDBConfig
}

var config *Config

func Init() {
	err := godotenv.Load("conf/.env")
	if err != nil {
		log.Printf("[Init] error on loading env from file: %+v", err)
	}

	config = &Config{
		AppName:    os.Getenv("APP_NAME"),
		AppAddress: os.Getenv("APP_ADDRESS"),
		MongoDBConfig: MongoDBConfig{
			DBName:    os.Getenv("MONGODB_DB_NAME"),
			DBAddress: os.Getenv("MONGODB_ADDRESS"),
		},
	}

	if config.AppName == "" {
		log.Panicf("[Init] app name cannot be empty")
	}

	if config.AppAddress == "" {
		log.Panicf("[Init] app address cannot be empty")
	}

	appEnv := Environment(os.Getenv("ENVIRONMENT"))

	if appEnv != EnvDev && appEnv != EnvProd {
		log.Panicf("[Init] app env must be dev or prod, found: %+v", appEnv)
	}

	config.Environment = appEnv

	if config.MongoDBConfig.DBAddress == "" ||
		config.MongoDBConfig.DBName == "" {
		log.Panic("[Init] mongodb config cannot be empty")
	}
}

func Get() *Config {
	return config
}
