package db

import (
	"hyneo-backend/graph/model"
	"hyneo-backend/internal/item"
	"hyneo-backend/pkg/logging"
	"hyneo-backend/pkg/mysql"
)

type repository struct {
	client mysql.Client
	logger *logging.Logger
}

func NewRepository(client mysql.Client, logger *logging.Logger) item.Repository {
	client.DB.AutoMigrate(model.Item{})
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) Create(item model.Item) error {
	r.logger.Tracef("Create item %s", item.Name)
	err := r.client.DB.Create(item).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(item model.Item) error {
	r.logger.Tracef("Update item %d", item.ID)
	err := r.client.DB.Save(item).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(item model.Item) error {
	r.logger.Tracef("Delete item %d", item.ID)
	err := r.client.DB.Delete(item).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetOne(id string) (model.Item, error) {
	r.logger.Tracef("Find item by id %s", id)
	var itm model.Item
	err := r.client.DB.Model(model.Item{}).First(&itm, id).Error
	if err != nil {
		return model.Item{}, err
	}
	return itm, nil
}

func (r *repository) GetAll() ([]*model.Item, error) {
	r.logger.Tracef("Find items")
	items := make([]*model.Item, 1)
	err := r.client.DB.Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}
