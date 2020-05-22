package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/AlexKoppara/Demeter/graph/generated"
	"github.com/AlexKoppara/Demeter/graph/model"
)

func (r *entityResolver) FindMenuItemByID(ctx context.Context, id string) (*model.MenuItem, error) {
	var menuItem *model.MenuItem
	switch id {
	case "menu_item_1":
		menuItem = &model.MenuItem{
			ID:           "menu_item_1",
			Name:         "chicken wings",
			Keyword:      "wings",
			PriceInCents: 1500,
		}
	case "menu_item_2":
		menuItem = &model.MenuItem{
			ID:           "menu_item_2",
			Name:         "french fries",
			Keyword:      "fries",
			PriceInCents: 500,
		}
	case "menu_item_3":
		menuItem = &model.MenuItem{
			ID:           "menu_item_3",
			Name:         "pizza",
			Keyword:      "pizza",
			PriceInCents: 2000,
		}
	}
	menuItem.Restaurant = &model.Restaurant{
		ID:      "restaurant_" + id,
		Name:    "Testaurant",
		Addr1:   "123 Sesame St",
		City:    "Brooklyn",
		State:   "NY",
		Country: "USA",
		Zip:     "12345",
		IsLive:  true,
	}
	return menuItem, nil
}

func (r *entityResolver) FindRestaurantByID(ctx context.Context, id string) (*model.Restaurant, error) {
	restaurant := &model.Restaurant{
		ID:      "restaurant_" + id,
		Name:    "Testaurant",
		Addr1:   "123 Sesame St",
		City:    "Brooklyn",
		State:   "NY",
		Country: "USA",
		Zip:     "12345",
		IsLive:  true,
	}
	return restaurant, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
