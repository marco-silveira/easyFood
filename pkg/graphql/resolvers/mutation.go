package graphql

import (
	"context"
	"errors"

	"easyfood/pkg/entity"
	"easyfood/pkg/graphql/gqlgen"
	"easyfood/pkg/graphql/models"
	"easyfood/services"
)

type mutationResolver struct{
	services services.All
}

func NewMutationResolver(s services.All) gqlgen.MutationResolver {
	return mutationResolver{services: s}
}

func (m mutationResolver) CreateDish(ctx context.Context, input models.CreateDishInput) (*models.Dish, error) {
	if input.Name == "" {
		return nil, errors.New("invalid dish name")
	}

	if input.Restaurant == 0 {
		return nil, errors.New("must be associated to a restaurant")
	}

	dish := entity.Dish{
		RestaurantID: input.Restaurant,
		CategoryID:   input.Category,
		Name:         input.Name,
		Price:        input.Price,
		CookTime:     input.CookTime,
	}

	err := m.services.Dish.Create(ctx, &dish)
	if err != nil {
		return nil, err
	}

	return models.NewDish(&dish)[0], nil
}

func (m mutationResolver) CreateUser(ctx context.Context, input models.CreateUserInput) (*models.User, error) {
	return models.NewUser(entity.User{}), nil
}

func (m mutationResolver) CreateCategory(ctx context.Context, name string) (bool, error) {
	return true, nil
}

func (m mutationResolver) CreateRestaurant(ctx context.Context, input models.CreateRestaurantInput) (*models.Restaurant, error) {
	if input.Name == "" {
		return nil, errors.New("invalid name")
	}

	if input.Address == "" {
		return nil, errors.New("invalid address")
	}

	if len(input.OpenDays) == 0 {
		return nil, errors.New("must specify open days")
	}

	if input.PhoneNumber == "" {
		return nil, errors.New("invalid phone number")
	}

	restaurant := entity.Restaurant{
		OpenHour:    input.OpenHour,
		CloseHour:   input.CloseHour,
		OpenDays:    models.GetEntityWeekdays(input.OpenDays),
		Name:        input.Name,
		Description: input.Description,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
	}

	err := m.services.Restaurant.Create(ctx, &restaurant)
	if err != nil {
		return nil, err
	}

	return models.NewRestaurant(&restaurant)[0], nil
}
