package graph

import (
	"hyneo-backend/internal/category"
	"hyneo-backend/internal/item"
	"hyneo-backend/internal/promo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ItemRepository     item.Repository
	CategoryRepository category.Repository
	PromoRepository    promo.Repository
}
