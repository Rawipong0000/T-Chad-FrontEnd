package model

import (
	"time"
)

type Product struct {
	Product_ID      int       `json:"product_id"`
	Name            string    `json:"product_name"`
	Product_User_ID int       `json:"product_user_id"`
	User_User_ID    int       `json:"user_user_id"`
	U_name          string    `json:"name"`
	Shopname        *string   `json:"shopname"`
	Price           float32   `json:"price"`
	Description     *string   `json:"description"`
	Size            string    `json:"size"`
	Img             string    `json:"img"`
	Selling         bool      `json:"selling"`
	Create_date     time.Time `json:"create_date"`
	Update_date     time.Time `json:"update_date"`
	Delflag         bool      `json:"delflag"`
}

type CartProduct struct {
	Product_ID      int       `json:"product_id" db:"product_id"`
	Name            string    `json:"product_name" db:"product_name"`
	Product_User_ID int       `json:"product_user_id" db:"product_user_id"`
	User_User_ID    int       `json:"user_user_id" db:"user_user_id"`
	U_name          string    `json:"name" db:"name"`
	Shopname        *string   `json:"shopname" db:"shopname"`
	Price           float32   `json:"price" db:"price"`
	Description     *string   `json:"description,omitempty" db:"description"`
	Size            string    `json:"size" db:"size"`
	Img             string    `json:"img" db:"img"`
	Selling         bool      `json:"selling" db:"selling"`
	Create_date     time.Time `json:"create_date" db:"create_date"`
	Update_date     time.Time `json:"update_date" db:"update_date"`
	Delflag         bool      `json:"delflag" db:"delflag"`
}
