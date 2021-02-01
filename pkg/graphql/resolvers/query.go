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

func (q queryResolver) Restaurant(ctx context.Context, category []string) ([]*models.Restaurant, error) {
	restaurant := models.NewRestaurant()
	return restaurant, nil
}

func (q queryResolver) User(ctx context.Context, id int) (*models.User, error) {
	u, err := q.services.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.NewUser(*u), nil
}
