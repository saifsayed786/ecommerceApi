package models

type Order struct {
	Id           int `json:"id,string"`
	CustomerId   int
	User         User   `gorm:"foreignKey:CustomerId" json:"-"`
	OrderAddress string `json:"orderAddress"`
	OrderDate    string `json:"orderDate"`
	OrderStatus  string `json:"orderStatus"`
}

type OrderDetail struct {
	Id        int `json:"id,string"`
	OrderId   int
	ProductId int
	Order     Order   `gorm:"foreignKey:OrderId" json:"-"`
	Product   Product `gorm:"foreignKey:ProductId" json:"-"`
	Price     int
	Quantity  int
}
