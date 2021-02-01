package graphql

import (
	"context"
	"easyfood/pkg/entity"
	"easyfood/pkg/graphql/gqlgen"
	"easyfood/pkg/graphql/models"
)

type mutationResolver struct{}

func NewMutationResolver() gqlgen.MutationResolver {
	return new(mutationResolver)
}

func (m mutationResolver) CreateUser(ctx context.Context, input models.CreateUserInput) (*models.User, error) {
	return models.NewUser(entity.User{}), nil
}
