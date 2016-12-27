package main

import (
	"fmt"

	"log"

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

//FindOrganizationByID retrieves the organization from the database
func (database *Database) FindOrganizationByID(organization *Organization, id int) error {
	database.DB.Preload("Users").
		Preload("Activities").
		First(organization, id)
	return nil
}

//ListRestaurants - database wrapper
func (database *Database) ListRestaurants(restaurants *[]Restaurant) error {
	database.DB.Preload("Menus").
		Preload("Activities").
		Find(restaurants)
	return nil
}

//CreateRestaurant - create Restaurant
func (database *Database) CreateRestaurant(restaurant *Restaurant) error {
	if err := database.DB.Create(restaurant).Error; err != nil {
		log.Printf("Failed to create Restaurant : %s", err)
		return fmt.Errorf("Failed to create Restaurant")
	}
	return nil
}

//UpdateRestaurant - update Restaurant
func (database *Database) UpdateRestaurant(restaurant *Restaurant) error {
	tx := database.DB.Begin()
	var restaurantDAO Restaurant
	if err := database.DB.First(&restaurantDAO, restaurant.ID).Error; err != nil {
		log.Printf("Record does not exists : %s", err)
		tx.Rollback()
		return fmt.Errorf("Record does not exists")
	}

	restaurantDAO.Name = restaurant.Name
	if err := database.DB.Save(restaurantDAO).Error; err != nil {
		log.Printf("Failed to Save Data : %s", err)
		tx.Rollback()
		return fmt.Errorf("Failed to Update Restaurant Data")
	}
	tx.Commit()
	return nil
}

//GetRestaurantByID - Restaurant by ID
func (database *Database) GetRestaurantByID(restaurant *Restaurant, id int) error {
	database.DB.Preload("Menus").
		Preload("Activities").
		First(restaurant, id)
	return nil
}

//UpdateMenu - update menu
func (database *Database) UpdateMenu(menu *Menu) error {
	var dbMenu Menu
	err := database.DB.Model(&dbMenu).
		Where("id = ?", menu.ID).
		Updates(map[string]interface{}{"name": menu.Name, "price": menu.Price}).
		Error
	if err != nil {
		return fmt.Errorf("Failed to update Menu: %s", err)
	}
	return nil
}

//DeleteMenu - delete menu
func (database *Database) DeleteMenu(menu *Menu) error {
	err := database.DB.Delete(Menu{}, "ID = ?", menu.ID).Error
	if err != nil {
		return fmt.Errorf("Failed to delete menu: %s", err)
	}
	return nil
}
