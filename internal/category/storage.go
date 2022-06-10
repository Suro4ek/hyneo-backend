package category

import "hyneo-backend/graph/model"

type Repository interface {
	Create(category model.Category) error
	Update(category model.Category) error
	Delete(category model.Category) error
	GetOne(id string) (model.Category, error)
	GetAll() ([]*model.Category, error)
}
