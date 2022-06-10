package model

type Category struct {
	ID    uint    `json:"id" gorm:"primarykey"`
	Name  string  `json:"name"`
	Items []*Item `json:"items" gorm:"foreignKey:CategoryID"`
}
