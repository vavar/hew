package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

//Database type
type Database struct {
	DB *gorm.DB
}

// InitDBConnection initiates a connection to the database and migrates database schema
func InitDBConnection() *Database {
	host := viper.GetString("db.host")
	user := viper.GetString("db.user")
	dbname := viper.GetString("db.dbname")
	dbConfig := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", host, user, dbname)
	db, err := gorm.Open("postgres", dbConfig)
	if err != nil {
		panic(fmt.Errorf("Failed to initiate a connection to the database: %s", err))
	}

	fmt.Println("Migrating database")
	db.AutoMigrate(&User{}, &Organization{}, &Restaurant{}, &Menu{}, &Activity{}, &OrderItem{})

	return &Database{db}
}

//GetOrganizationByID retrieves the organization from the database
func (database *Database) GetOrganizationByID(organization *Organization, id int) {
	database.DB.First(organization, id)
}

//ListRestaurants - database wrapper
func (database *Database) ListRestaurants() ([]Restaurant, error) {
	return []Restaurant{}, nil
}

//FindRestaurantByID - Restaurant by ID
func (database *Database) FindRestaurantByID(id string) (Restaurant, error) {
	return Restaurant{}, nil
}
