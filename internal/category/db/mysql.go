package db

import (
	"fmt"
	"hyneo-backend/graph/model"
	"hyneo-backend/internal/category"
	"hyneo-backend/pkg/logging"
	"hyneo-backend/pkg/mysql"
)

type repository struct {
	client mysql.Client
	logger *logging.Logger
}

func NewRepository(client mysql.Client, logger *logging.Logger) category.Repository {
	client.DB.AutoMigrate(model.Category{})
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) Create(category model.Category) error {
	r.logger.Tracef("Create category %s", category.Name)
	err := r.client.DB.Create(category).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(category model.Category) error {
	r.logger.Tracef("Update category %s", category.ID)
	err := r.client.DB.Save(category).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(category model.Category) error {
	r.logger.Tracef("Delete category %d", category.ID)
	err := r.client.DB.Delete(category).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetOne(id string) (model.Category, error) {
	r.logger.Tracef("Find category by id %s", id)
	var cat model.Category
	err := r.client.DB.Model(model.Category{}).First(&cat, id).Error
	if err != nil {
		return model.Category{}, err
	}
	return cat, nil
}

func (r *repository) GetAll() ([]*model.Category, error) {
	r.logger.Tracef("Find categories")
	categories := make([]*model.Category, 1)
	err := r.client.DB.Model(model.Category{}).Preload("Items").Find(&categories).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(categories[0])
	return categories, nil
}
