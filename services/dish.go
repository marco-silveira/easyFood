package services

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"easyfood/pkg/entity"
)

type DishService interface {
	Get(ctx context.Context, id *int) ([]*entity.Dish, error)
	GetByCategory(ctx context.Context, categoryID int) ([]*entity.Dish, error)
	GetByRestaurant(ctx context.Context, restaurantID int) ([]*entity.Dish, error)
}

type dishService struct {
	db *sqlx.DB
}

func NewDishService(db *sqlx.DB) DishService {
	return dishService{db: db}
}

func (d dishService) Get(ctx context.Context, id *int) ([]*entity.Dish, error) {
	result := make([]*entity.Dish, 0)
	query := `
		SELECT * FROM pratos
	`
	if id != nil {
		query = fmt.Sprintf("SELECT * FROM pratos WHERE id = %d", *id)
	}

	err := d.db.Select(&result, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d dishService) GetByCategory(ctx context.Context, categoryID int) ([]*entity.Dish, error) {
	result := make([]*entity.Dish, 0)

	query := `
		SELECT * FROM pratos WHERE id_categoria = ?
	`

	err := d.db.Select(&result, query, categoryID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d dishService) GetByRestaurant(ctx context.Context, restaurantID int) ([]*entity.Dish, error) {
	result := make([]*entity.Dish, 0)

	query := `
		SELECT * FROM pratos WHERE id_restaurante = ?
	`

	err := d.db.Select(&result, query, restaurantID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
