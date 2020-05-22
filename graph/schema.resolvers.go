package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/AlexKoppara/Demeter/graph/generated"
	"github.com/AlexKoppara/Demeter/graph/model"
)

func (r *queryResolver) Restaurant(ctx context.Context) (*model.Restaurant, error) {
	restaurant := &model.Restaurant{
		ID:      "restaurant_1",
		Name:    "Testaurant",
		Addr1:   "123 Sesame St",
		City:    "Brooklyn",
		State:   "NY",
		Country: "USA",
		Zip:     "12345",
		IsLive:  true,
	}

	var menuItems []*model.MenuItem
	menuItems = append(menuItems, &model.MenuItem{
		ID:           "menu_item_1",
		Restaurant:   restaurant,
		Name:         "chicken wings",
		Keyword:      "wings",
		PriceInCents: 1500,
	})
	menuItems = append(menuItems, &model.MenuItem{
		ID:           "menu_item_2",
		Restaurant:   restaurant,
		Name:         "french fries",
		Keyword:      "fries",
		PriceInCents: 500,
	})
	menuItems = append(menuItems, &model.MenuItem{
		ID:           "menu_item_3",
		Restaurant:   restaurant,
		Name:         "pizza",
		Keyword:      "pizza",
		PriceInCents: 2000,
	})

	restaurant.MeuItems = menuItems
	return restaurant, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
