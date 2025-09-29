package model

import "time"

type Purchasing struct {
	Purc_id       int       `json:"purchasing_id"`
	Tran_id       int       `json:"transaction_id"`
	Product_ID    int       `json:"product_id"`
	Discount_code string    `json:"discount_code"`
	Tracking      string    `json:"tracking"`
	Create_date   time.Time `json:"create_date"`
	Update_date   time.Time `json:"update_date"`
	Delflag       bool      `json:"delflag"`
}

type Transaction struct {
	Tran_id     int       `json:"transaction_id"`
	User_ID     int       `json:"user_id"`
	Discount    float32   `json:"discount"`
	Total       float32   `json:"total"`
	Status_code int       `json:"status_code"`
	Address     string    `json:"address"`
	Create_date time.Time `json:"create_date"`
	Update_date time.Time `json:"update_date"`
	Delflag     bool      `json:"delflag"`
}

type Sub_Transaction struct {
	Sub_Tran_id   int       `json:"sub_tran_id"`
	Tran_id       int       `json:"tran_id"`
	Seller_ID     int       `json:"seller_id"`
	Discount_code string    `json:"discount_code"`
	Sub_Total     float32   `json:"sub_total"`
	Tracking      string    `json:"tracking"`
	Status_code   int       `json:"status_code"`
	Create_date   time.Time `json:"create_date"`
	Update_date   time.Time `json:"update_date"`
	Delflag       bool      `json:"delflag"`
}

type Myshop_ordering struct {
	Sub_Tran_id         int       `json:"sub_tran_id"`
	Sub_Tran_Tran_id    int       `json:"tran_id"`
	Transaction_Tran_id int       `json:"transaction_id"`
	Transaction_User_ID int       `json:"tranaction_user_id"`
	User_User_ID        int       `json:"user_user_id"`
	Name                string    `json:"Name"`
	Address             *string   `json:"address"`
	Shopname            string    `json:"shopname"`
	Purchase_Tran_id    int       `json:"purchase_transaction_id"`
	Purchase_Product_ID int       `json:"purchase_product_id"`
	Product_Product_ID  int       `json:"product_product_id"`
	Product_User_ID     int       `json:"product_user_id"`
	Product_Name        string    `json:"product_name"`
	Discount_code       *string   `json:"discount_code"`
	Tracking            *string   `json:"tracking"`
	Sub_Total           float32   `json:"sub_total"`
	Sub_Status_code     int       `json:"sub_status_code"`
	Status_Status_code  int       `json:"status_code"`
	Status_name         string    `json:"status_name"`
	Color               string    `json:"color"`
	Create_date         time.Time `json:"create_date"`
	Update_date         time.Time `json:"update_date"`
	Delflag             bool      `json:"delflag"`
}

type History_ordering struct {
	Sub_Tran_id         int       `json:"sub_tran_id"`
	Sub_Tran_Tran_id    int       `json:"tran_id"`
	Transaction_Tran_id int       `json:"transaction_id"`
	Sub_Tran_Seller_id  int       `json:"seller_id"`
	User_User_ID        int       `json:"user_user_id"`
	Name                string    `json:"name"`
	Address             *string   `json:"address"`
	Shopname            *string   `json:"shopname"`
	Purchase_Tran_id    int       `json:"purchase_transaction_id"`
	Purchase_Product_ID int       `json:"purchase_product_id"`
	Product_Product_ID  int       `json:"product_product_id"`
	Product_Name        string    `json:"product_name"`
	Product_User_ID     int       `json:"product_user_id"` // << ต้องมีฟิลด์นี้
	Discount_code       *string   `json:"discount_code"`
	Tracking            *string   `json:"tracking"`
	Sub_Total           float32   `json:"sub_total"`
	Sub_Status_code     int       `json:"sub_status_code"`
	Status_Status_code  int       `json:"status_code"`
	Status_name         string    `json:"status_name"`
	Color               string    `json:"color"`
	Create_date         time.Time `json:"create_date"`
	Update_date         time.Time `json:"update_date"`
	Delflag             bool      `json:"delflag"`
}
