package main

import (
	"fmt"
	"reflect"
	"time"

	"log"

	"github.com/jinzhu/gorm"
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

//Create - Create an Object
func (database *Database) Create(object interface{}) error {
	if err := database.DB.Create(object).Error; err != nil {
		return fmt.Errorf("Failed to create a '%s': %s", reflect.TypeOf(object), err)
	}
	return nil
}

//Update - Update an organization
func (database *Database) Update(object interface{}) error {
	if err := database.DB.Model(object).Updates(object).Error; err != nil {
		return fmt.Errorf("Failed to update '%s': %s", reflect.TypeOf(object), err)
	}
	return nil
}

//ListUsers - List all users
func (database *Database) ListUsers(users *[]User) error {
	err := database.DB.Debug().Select("id, username, email,organization_id").Find(users).Error
	if err != nil {
		return fmt.Errorf("Failed to select all users: %s", err)
	}

	return nil
}

//FindUserByID - Get a user with the given ID
func (database *Database) FindUserByID(user *User, id int) error {
	err := database.DB.Preload("OrderItems").
		Where("id = ?", id).
		First(user, id).Error

	if err != nil {
		return fmt.Errorf("Failed to get a user with ID = %d: %s", id, err)
	}

	return nil
}

//FindUserByEmail - Get a user with the given Email
func (database *Database) FindUserByEmail(user *User, email string) error {
	err := database.DB.Where("email = ?", email).First(user).Error
	if err != nil {
		log.Printf("Failed to get a user with e-mail = %s: %s", email, err)
		return fmt.Errorf("Failed to get a user with e-mail = %s: %s", email, err)
	}

	return nil
}

//ListOrganizations - List all organizations
func (database *Database) ListOrganizations(organizations *[]Organization) error {
	err := database.DB.Preload("Users").
		Preload("Activities").
		Find(organizations).Error

	if err != nil {
		return fmt.Errorf("Failed to select all organizations: %s", err)
	}

	return nil
}

//FindOrganizationByID retrieves the organization from the database
func (database *Database) FindOrganizationByID(organization *Organization, id int) error {
	err := database.DB.Preload("Users").
		Preload("Activities").
		Where("id = ?", id).
		First(organization, id).Error

	if err != nil {
		return fmt.Errorf("Failed to find an organization with ID = %d: %s", id, err)
	}

	return nil
}

//ListRestaurants - database wrapper
func (database *Database) ListRestaurants(restaurants *[]Restaurant) error {
	err := database.DB.Preload("Menus").
		Find(restaurants).Error

	if err != nil {
		return fmt.Errorf("Failed to select all restaurants: %s", err)
	}

	return nil
}

//CreateRestaurant - create Restaurant
func (database *Database) CreateRestaurant(restaurant *Restaurant) error {
	if err := database.DB.Create(restaurant).Error; err != nil {
		log.Printf("Failed to create Restaurant: %s", err)
		return fmt.Errorf("Failed to create Restaurant")
	}
	return nil
}

//UpdateRestaurant - update Restaurant
func (database *Database) UpdateRestaurant(restaurant *Restaurant) error {
	tx := database.DB.Begin()
	var restaurantDAO Restaurant
	if err := database.DB.First(&restaurantDAO, restaurant.ID).Error; err != nil {
		log.Printf("Record does not exists: %s", err)
		tx.Rollback()
		return fmt.Errorf("Record does not exists")
	}

	restaurantDAO.Name = restaurant.Name
	if err := database.DB.Save(restaurantDAO).Error; err != nil {
		log.Printf("Failed to Save Data: %s", err)
		tx.Rollback()
		return fmt.Errorf("Failed to Update Restaurant Data")
	}
	tx.Commit()
	return nil
}

//FindRestaurantByID - Restaurant by ID
func (database *Database) FindRestaurantByID(restaurant *Restaurant, id int) error {
	err := database.DB.Preload("Menus").
		Where("id = ?", id).
		First(restaurant, id).Error

	if err != nil {
		return fmt.Errorf("Failed to find a restaurant with ID = %d: %s", id, err)
	}

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

//ListActivities - List all activities made by an organization
func (database *Database) ListActivities(activities *[]Activity, organizationID int, status string) error {
	now := time.Now().UTC().Format(time.UnixDate)

	filteredActivities := database.DB
	switch status {
	case "open":
		filteredActivities = database.DB.Where("closed_at > ?", now)
	case "closed":
		filteredActivities = database.DB.Where("closed_at <= ?", now)
	}

	err := filteredActivities.Order("closed_at asc").
		Preload("OrderItems").
		Preload("Restaurants").
		Preload("Restaurants.Menus").
		Where(&Activity{OrganizationID: organizationID}).
		Find(activities).Error

	if err != nil {
		return fmt.Errorf("Failed to select all activites: %s", err)
	}

	return nil
}

//FindActivityByID - Get an activity with the given ID
func (database *Database) FindActivityByID(activity *Activity, id int) error {
	err := database.DB.Preload("OrderItems").
		Preload("Restaurants").
		Preload("Restaurants.Menus").
		Where("id = ?", id).
		First(activity).Error

	if err != nil {
		return fmt.Errorf("Failed to get an activity with ID = %d: %s", id, err)
	}

	return nil
}

//AddActivityRestaurant - add Activity Restaurant
func (database *Database) AddActivityRestaurant(activity *Activity, id int) error {
	var restaurant Restaurant
	if notFound := database.FindRestaurantByID(&restaurant, id); notFound != nil {
		return notFound
	}
	if err := database.DB.Model(activity).Association("Restaurants").Append(restaurant).Error; err != nil {
		return fmt.Errorf("Failed to relation an activity with restaurant ID = %d: %s", id, err)
	}
	database.FindActivityByID(activity, activity.ID)
	return nil
}

//DeleteActivityRestaurant - Delete a row from table activities_restaurants
func (database *Database) DeleteActivityRestaurant(activity *Activity, restaurantID int) error {
	var restaurant Restaurant
	if notFound := database.FindRestaurantByID(&restaurant, restaurantID); notFound != nil {
		return notFound
	}
	if err := database.DB.Model(activity).Association("Restaurants").Delete(restaurant).Error; err != nil {
		return fmt.Errorf("Failed to remove the relation between the activity %d and the restaurant %d: %s", activity.ID, restaurantID, err)
	}
	database.FindActivityByID(activity, activity.ID)
	return nil
}

//ListOrderItems - List all order items made by a user
func (database *Database) ListOrderItems(items *[]OrderItem, userID int) error {
	var user = &User{}
	user.ID = userID
	if err := database.DB.Debug().Model(&user).Related(items).Error; err != nil {
		return fmt.Errorf("Failed to select all order items ")
	}
	return nil
}

//FindOrderItemByID - Get an order item with the given ID
func (database *Database) FindOrderItemByID(item *OrderItem, id int) error {
	if err := database.DB.First(item, id).Error; err != nil {
		return fmt.Errorf("Failed to get an order item with ID = %d: %s", id, err)
	}
	return nil
}

//FindOrderItemsByUserID
func (database *Database) FindOrderItemsByUserID(items *[]OrderHistory, id int) error {
	rows, err := database.DB.
		Raw("select menus.id, menus.name, menus.price, activities.name, users.username, order_items.created_at "+
			"from order_items,menus,users,activities "+
			"where "+
			"order_items.user_id = users.id "+
			"AND order_items.menu_id = menus.id "+
			"AND order_items.activity_id = activities.id "+
			"AND order_items.user_id = ?", id).
		Rows()
	defer rows.Close()
	for rows.Next() {

		item := OrderHistory{}
		rows.Scan(&item.MenuID,
			&item.MenuName,
			&item.MenuPrice,
			&item.ActivityName,
			&item.Username,
			&item.Date)
		*items = append(*items, item)
	}
	return err
}
