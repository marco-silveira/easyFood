package services

import "github.com/jmoiron/sqlx"

type All struct {
	User UserService
	Dish DishService
}

func NewServices(db *sqlx.DB) All {
	return All{
		User: NewUserService(db),
		Dish: NewDishService(db),
	}
}
