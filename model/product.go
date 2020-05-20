package model


type Product struct {
	ID int `gorm:"primary_key" json:"id"`
	Name string `json: "name"`
	Info string `json: "info"`
	Url string `json: "url"`
	Price float64 `json: "price"`
}

func (Product) TableName() string{ return "products"}
