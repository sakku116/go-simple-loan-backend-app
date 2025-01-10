package config

import (
	"github.com/spf13/viper"
)

type EnvsSchema struct {
	HOST                   string
	PORT                   int
	LOG_LEVEL              string
	JWT_SECRET_KEY         string
	JWT_EXP_MINS           int
	REFRESH_TOKEN_EXP_MINS int

	MYSQL_HOST       string
	MYSQL_PORT       int
	MYSQL_USER       string
	MYSQL_PASSWORD   string
	MYSQL_DATABASE   string
	MINIO_ACCESS_KEY string
	MINIO_SECRET_KEY string

	INITIAL_USER_USERNAME  string
	INITIAL_USER_PASSWORD  string
	INITIAL_ADMIN_USERNAME string
	INITIAL_ADMIN_PASSWORD string
}

var Envs *EnvsSchema

func envInitiator() {
	Envs = &EnvsSchema{
		HOST:                   viper.GetString("HOST"),
		PORT:                   viper.GetInt("PORT"),
		LOG_LEVEL:              viper.GetString("LOG_LEVEL"),
		JWT_SECRET_KEY:         viper.GetString("JWT_SECRET_KEY"),
		JWT_EXP_MINS:           viper.GetInt("JWT_EXP_MINS"),
		REFRESH_TOKEN_EXP_MINS: viper.GetInt("REFRESH_TOKEN_EXP_MINS"),

		MYSQL_HOST:       viper.GetString("MYSQL_HOST"),
		MYSQL_PORT:       viper.GetInt("MYSQL_PORT"),
		MYSQL_USER:       viper.GetString("MYSQL_USER"),
		MYSQL_PASSWORD:   viper.GetString("MYSQL_PASSWORD"),
		MYSQL_DATABASE:   viper.GetString("MYSQL_DATABASE"),
		MINIO_ACCESS_KEY: viper.GetString("MINIO_ACCESS_KEY"),
		MINIO_SECRET_KEY: viper.GetString("MINIO_SECRET_KEY"),

		INITIAL_USER_USERNAME:  viper.GetString("INITIAL_USER_USERNAME"),
		INITIAL_USER_PASSWORD:  viper.GetString("INITIAL_USER_PASSWORD"),
		INITIAL_ADMIN_USERNAME: viper.GetString("INITIAL_ADMIN_USERNAME"),
		INITIAL_ADMIN_PASSWORD: viper.GetString("INITIAL_ADMIN_PASSWORD"),
	}
}

func InitEnv(filepath string) {
	viper.SetConfigType("env")
	viper.SetConfigFile(filepath)
	if err := viper.ReadInConfig(); err != nil {
		logger.Warningf("error loading environment variables from %s: %w", filepath, err)
	}
	viper.AutomaticEnv()
	envInitiator()
}
