package db

import (
	"hyneo-backend/graph/model"
	"hyneo-backend/internal/promo"
	"hyneo-backend/pkg/logging"
	"hyneo-backend/pkg/mysql"
)

type repository struct {
	client mysql.Client
	logger *logging.Logger
}

func NewRepository(client mysql.Client, logger *logging.Logger) promo.Repository {
	client.DB.AutoMigrate(&model.PromoCode{})
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) Create(promo model.PromoCode) error {
	r.logger.Tracef("Create new promo ")
	err := r.client.DB.Create(promo).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(promo model.PromoCode) error {
	r.logger.Tracef("Update promo %d", promo.ID)
	err := r.client.DB.Save(promo).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(promo model.PromoCode) error {
	r.logger.Tracef("delete promo %d", promo.ID)
	err := r.client.DB.Delete(promo).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CheckPromo(name string) bool {
	r.logger.Tracef("Find promo by name %s", name)
	var pro model.PromoCode
	err := r.client.DB.Model(model.PromoCode{}).First(&pro, "name = ?", name).Error
	if err != nil {
		return false
	}
	return true
}

func (r *repository) GetOne(id string) (model.PromoCode, error) {
	r.logger.Tracef("Find promo by id %s", id)
	var pro model.PromoCode
	err := r.client.DB.Model(model.PromoCode{}).First(&pro, id).Error
	if err != nil {
		return model.PromoCode{}, err
	}
	return pro, nil
}

func (r *repository) GetAll() ([]*model.PromoCode, error) {
	r.logger.Tracef("Find promocodes")
	promocodies := make([]*model.PromoCode, 1)
	err := r.client.DB.Find(&promocodies).Error
	if err != nil {
		return nil, err
	}
	return promocodies, nil
}
