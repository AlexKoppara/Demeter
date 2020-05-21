package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/AlexKoppara/Demeter/graph/generated"
	"github.com/AlexKoppara/Demeter/graph/model"
)

func (r *entityResolver) FindRestaurantByID(ctx context.Context, id string) (*model.Restaurant, error) {
	restaurant := &model.Restaurant{
		ID:      "restaurant_1",
		Name:    "Testaurant",
		Addr1:   "123 Sesame St",
		City:    "Brooklyn",
		State:   "NY",
		Country: "USA",
		Zip:     "12345",
		IsLive:  true,
		Menu: &model.Menu{
			ID: "menu_1",
		},
	}
	return restaurant, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
