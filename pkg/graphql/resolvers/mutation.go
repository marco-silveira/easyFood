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

func (m mutationResolver) CreateDish(ctx context.Context, input models.CreateDishInput) (bool, error) {
	return true, nil
}

func (m mutationResolver) CreateUser(ctx context.Context, input models.CreateUserInput) (*models.User, error) {
	if input.Email == "" {
		return nil, errors.New("invalid email")
	}

	if input.Senha == "" {
		return nil, errors.New("invalid password")
	}

	user := entity.User{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    input.Senha,
	}

	err := m.services.User.Create(ctx, &user)
	if err != nil {
		return nil, err
	}
	return models.NewUser(user), nil
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
