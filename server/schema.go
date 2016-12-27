package main

import "time"

//User record
type User struct {
	ID             int         `json:"id"                        gorm:"primary_key;AUTO_INCREMENT"`
	Username       string      `json:"username"                  gorm:"type:varchar(200)"`
	Email          string      `json:"email,omitempty"           gorm:"type:varchar(200)"`
	Password       string      `json:"password,omitempty"        gorm:"type:varchar(100)"`
	OrganizationID int         `json:"organization_id,omitempty"`
	OrderItems     []OrderItem `json:"order_items"`
}

//Organization record
type Organization struct {
	ID         int        `json:"id"         gorm:"primary_key;AUTO_INCREMENT"`
	Name       string     `json:"name"       gorm:"type:varchar(200)"`
	Users      []User     `json:"users"`
	Activities []Activity `json:"activities"`
}

//Restaurant record
type Restaurant struct {
	ID         int        `json:"id"         gorm:"primary_key;AUTO_INCREMENT"`
	Name       string     `json:"name"       gorm:"type:varchar(500)"`
	Menus      []Menu     `json:"menus"`
	Activities []Activity `json:"activities" gorm:"many2many:restaurants_activities"`
}

//Menu record
type Menu struct {
	ID           int         `json:"id"           gorm:"primary_key;AUTO_INCREMENT"`
	Name         string      `json:"name"         gorm:"type:varchar(255)"`
	Price        float32     `json:"price"`
	RestaurantID int         `json:"restaurant_id"`
	OrderItems   []OrderItem `json:"order_items"`
}

//Activity record
type Activity struct {
	ID             int         `json:"id"              gorm:"primary_key;AUTO_INCREMENT"`
	Name           string      `json:"name"            gorm:"type:varchar(255)"`
	IsOpen         bool        `json:"is_open"`
	OrganizationID int         `json:"organization_id"`
	OrderItems     []OrderItem `json:"order_items"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
}

//OrderItem record
type OrderItem struct {
	ID         int       `json:"id"`
	ActivityID int       `json:"activity_id"`
	UserID     int       `json:"user_id"`
	MenuID     int       `json:"menu_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}