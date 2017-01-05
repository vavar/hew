package main

import "time"

//User record
type User struct {
	ID             int         `json:"id"              gorm:"primary_key;AUTO_INCREMENT"`
	Username       string      `json:"username"        gorm:"type:varchar(200)"`
	Email          string      `json:"email"           gorm:"type:varchar(200)"`
	Password       string      `json:"password,omitempty"        gorm:"type:varchar(100)"`
	Role           string      `json:"role,omitempty"            gorm:"type:varchar(20)"` // values: user, admin, system_admin
	OrganizationID int         `json:"organization_id" `
	OrderItems     []OrderItem `json:"order_items,omitempty" gorm:"ForeignKey:UserID"`
}

//Organization record
type Organization struct {
	ID         int        `json:"id"         gorm:"primary_key;AUTO_INCREMENT"`
	Name       string     `json:"name"       gorm:"type:varchar(200)"`
	Users      []User     `json:"users" gorm:"ForeignKey:OrganizationID"`
	Activities []Activity `json:"activities" gorm:"ForeignKey:OrganizationID"`
}

//Restaurant record
type Restaurant struct {
	ID    int    `json:"id"    gorm:"primary_key:true;AUTO_INCREMENT"`
	Name  string `json:"name"  gorm:"type:varchar(500)"`
	Phone string `json:"phone" gorm:"type:varchar(50)"`
	Menus []Menu `json:"menus" gorm:"ForeignKey:RestaurantID"`
}

//Menu record
type Menu struct {
	ID           int         `json:"id"           gorm:"primary_key;AUTO_INCREMENT"`
	Name         string      `json:"name"         gorm:"type:varchar(255)"`
	Price        float32     `json:"price"`
	RestaurantID int         `json:"restaurant_id"`
	OrderItems   []OrderItem `json:"order_items" gorm:"ForeignKey:MenuID"`
}

//Activity record
type Activity struct {
	ID             int          `json:"id"              gorm:"primary_key:true;AUTO_INCREMENT"`
	Name           string       `json:"name"            gorm:"type:varchar(255)"`
	ClosedAt       time.Time    `json:"closed_at"`
	OrganizationID int          `json:"organization_id"`
	OrderItems     []OrderItem  `json:"order_items" gorm:"ForeignKey:ActivityID"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	Restaurants    []Restaurant `json:"restaurants"     gorm:"many2many:activities_restaurants"`
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

//OrderHistory View
type OrderHistory struct {
	MenuID       int       `json:"menu_id"`
	MenuName     string    `json:"menu_name"`
	MenuPrice    float32   `json:"menu_price"`
	ActivityName string    `json:"activity_name"`
	Username     string    `json:"username"`
	Date         time.Time `json:"order_date"`
}
