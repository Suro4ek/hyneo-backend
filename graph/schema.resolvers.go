package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hyneo-backend/graph/generated"
	"hyneo-backend/graph/model"
	"time"

	"github.com/Tnze/go-mc/bot"
)

func (r *categoryResolver) ID(ctx context.Context, obj *model.Category) (int, error) {
	return int(obj.ID), nil
}

func (r *itemResolver) ID(ctx context.Context, obj *model.Item) (int, error) {
	return int(obj.ID), nil
}

func (r *mutationResolver) CheckPromo(ctx context.Context, name string) (*model.PromoCode, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Buy(ctx context.Context, buy model.BuyInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	items, err := r.ItemRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return items, err
}

func (r *queryResolver) Category(ctx context.Context) ([]*model.Category, error) {
	categories, err := r.CategoryRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return categories, err
}

func (r *queryResolver) Minecraft(ctx context.Context) (*model.Minecraft, error) {
	resp, delay, err := bot.PingAndList("mc.hyneo.ru")
	if err != nil {
		return nil, errors.New("error ping server")
	}
	var s status
	err = json.Unmarshal(resp, &s)
	if err != nil {
		return nil, errors.New("fail unmarshal json")
	}
	s.Delay = delay
	return &model.Minecraft{OnlinePlayers: s.Players.Online}, nil
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Item returns generated.ItemResolver implementation.
func (r *Resolver) Item() generated.ItemResolver { return &itemResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type itemResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type status struct {
	Players struct {
		Max    int
		Online int
	}
	Delay time.Duration
}
