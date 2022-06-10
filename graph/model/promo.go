package model

type PromoCode struct {
	ID       uint   `json:"-" gorm:"primarykey"`
	Name     string `json:"-"`
	Discount int    `json:"discount"`
}
