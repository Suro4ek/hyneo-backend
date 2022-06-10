package promo

import "hyneo-backend/graph/model"

type Repository interface {
	Create(promo model.PromoCode) error
	Update(promo model.PromoCode) error
	Delete(promo model.PromoCode) error
	CheckPromo(name string) bool
	GetOne(id string) (model.PromoCode, error)
	GetAll() ([]*model.PromoCode, error)
}
