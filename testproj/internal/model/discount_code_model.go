package model

import (
	"time"
)

type Discount_code struct {
	Discount_ID         int       `json:"discount_id"`
	Seller_ID           int       `json:"seller_id"`
	Discount_code       string    `json:"discount_code"`
	Limit               int       `json:"limit"`
	Used                int       `json:"used"`
	Discount_by_percent float32   `json:"discount_by_percent"`
	Discount_by_number  float32   `json:"discount_by_number"`
	Minimum_total       float32   `json:"minimum_total"`
	Maximum_discount    float32   `json:"maximum_discount"`
	Create_date         time.Time `json:"create_date"`
	Update_date         time.Time `json:"update_date"`
	Delflag             bool      `json:"delflag"`
}
