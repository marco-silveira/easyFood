package graphql

import (
	"context"

	"easyfood/pkg/graphql/gqlgen"
	"easyfood/pkg/graphql/models"
	"easyfood/services"
)

type queryResolver struct {
	services services.All
}

func NewQueryResolver(services services.All) gqlgen.QueryResolver {
	return queryResolver{services: services}
}

func (q queryResolver) Category(ctx context.Context, id int) (*models.Category, error) {
	category := models.NewCategory()
	return category, nil
}

func (q queryResolver) Dish(ctx context.Context, id *int) ([]*models.Dish, error) {
	d, err := q.services.Dish.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.NewDish(d...), nil
}

func (q queryResolver) Restaurant(ctx context.Context, id int) (*models.Restaurant, error) {
	//TODO fetch restaurante
	restaurant := models.NewRestaurant()
	return restaurant[0], nil
}

func (q queryResolver) User(ctx context.Context, id int) (*models.User, error) {
	u, err := q.services.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.NewUser(*u), nil
}
