package main

import "time"

//User record
type User struct {
	ID           int           `json:"id"`
	Username     string        `json:"username"`
	Email        string        `json:"email,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
}

//Organization record
type Organization struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Restaurant record
type Restaurant struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Menus []*Menu `json:"menu"`
}

//Menu record
type Menu struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

//Activity in System
type Activity struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	CreateDate   time.Time     `json:"create_date"`
	UpdateDate   time.Time     `json:"update_date"`
	IsOpen       bool          `json:"IsOpen"`
	Organization *Organization `json:"organization"`
	Restaurants  []*Restaurant `json:"restaurants"`
	OrderItems   []*OrderItem  `json:"order_items"`
}

//OrderItem record
type OrderItem struct {
	ID         int       `json:"id"`
	User       *User     `json:"user"`
	Menu       *Menu     `json:"menu"`
	CreateDate time.Time `json:"create_date"`
	UpdateDate time.Time `json:"update_date"`
}
