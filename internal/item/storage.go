package item

import "hyneo-backend/graph/model"

type Repository interface {
	Create(item model.Item) error
	Update(item model.Item) error
	Delete(item model.Item) error
	GetOne(id string) (model.Item, error)
	GetAll() ([]*model.Item, error)
}
