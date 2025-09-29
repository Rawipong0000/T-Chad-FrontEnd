package model

import "time"

type Users struct {
	User_ID     int       `json:"user_id"`
	Name        string    `json:"Name"`
	Lastname    string    `json:"Last_name"`
	Email       string    `json:"Email"`
	Password    string    `json:"password"`
	Shopname    *string   `json:"shopname"`
	Phone       *string   `json:"phone"`
	Address     *string   `json:"address"`
	SubDistrict *string   `json:"subdistrict"`
	District    *string   `json:"district"`
	Province    *string   `json:"province"`
	Postal_code *string   `json:"postal_code"`
	Create_date time.Time `json:"create_date"`
	Update_date time.Time `json:"update_date"`
	Delflag     bool      `json:"delflag"`
}

type Province struct {
	Province_ID int     `json:"province_id"`
	NameTH      string  `json:"name_th"`
	NameEN      *string `json:"name_en"`
}

type District struct {
	District_ID int     `json:"district_id"`
	Province_ID int     `json:"province_id"`
	NameTH      string  `json:"name_th"`
	NameEN      *string `json:"name_en"`
}

type SubDistrict struct {
	SubDistrict_ID int     `json:"subdistrict_id"`
	District_ID    int     `json:"district_id"`
	NameTH         string  `json:"name_th"`
	NameEN         *string `json:"name_en"`
	Lat            *string `json:"lat"`
	Long           *string `json:"long"`
	Zipcode        string  `json:"zipcode"`
}
