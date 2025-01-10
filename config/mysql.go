package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySqlDB() *gorm.DB {
	var err error

	// connect to default database
	logger.Debugf("connecting to default database: mysql")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Envs.MYSQL_USER,
		Envs.MYSQL_PASSWORD,
		Envs.MYSQL_HOST,
		Envs.MYSQL_PORT,
		"mysql",
	)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("failed to connect to the database: %v", err)
	}

	// create specified database
	logger.Debugf("ensuring database: %s", Envs.MYSQL_DB)
	err = DB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", Envs.MYSQL_DB)).Error
	if err != nil {
		logger.Warningf("Failed to create database: %v", err)
	}

	// connect to specified database

	logger.Debugf("connecting to database: %s", Envs.MYSQL_DB)
	dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Envs.MYSQL_USER,
		Envs.MYSQL_PASSWORD,
		Envs.MYSQL_HOST,
		Envs.MYSQL_PORT,
		Envs.MYSQL_DB,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("failed to connect to the database: %v", err)
	}

	return DB
}
