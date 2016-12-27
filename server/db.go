package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

type Database struct {
	DB *gorm.DB
}

func InitDBConnection() *Database {
	host := viper.GetString("db.host")
	user := viper.GetString("db.user")
	dbname := viper.GetString("db.dbname")
	dbConfig := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", host, user, dbname)
	db, err := gorm.Open("postgres", dbConfig)
	if err != nil {
		panic(fmt.Errorf("Failed to initiate a connection to the database: %s", err))
	}
	defer db.Close()

	return &Database{db}
}
