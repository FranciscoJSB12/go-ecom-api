package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost        string
	Port              string
	DBUser            string
	MySQLRootPassword string
	DBAddress         string
	DBName            string
}

var Envs = initConfig()

func initConfig() Config {

	godotenv.Load()

	return Config{
		PublicHost:        getEnv("PUBLIC_HOST", "http://localhost"),
		Port:              getEnv("PORT", "8080"),
		DBUser:            getEnv("DB_USER", "root"),
		MySQLRootPassword: getEnv("MYSQL_ROOT_PASSWORD", "admin"),
		DBAddress:         fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		//default address > 127.0.0.1:3306
		DBName: getEnv("DB_NAME", "ecom"),
	}
}

func getEnv(key, fallback string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
