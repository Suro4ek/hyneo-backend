package model

import "time"

type Item struct {
	ID         uint `json:"id" gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string  `json:"name"`
	Price      int     `json:"price"`
	Desc       string  `json:"desc"`
	Img        *string `json:"img"`
	Discprice  int     `json:"discprice"`
	Command    string  `json:"-"`
	CategoryID uint    `json:"-"`
}
