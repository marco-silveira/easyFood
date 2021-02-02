package entity

type Dish struct {
	Id           int     `db:"id"`
	RestaurantID int     `db:"id_restaurante"`
	CategoryID   *int    `db:"id_categoria"`
	Name         string  `db:"nome"`
	Price        float64 `db:"preco"`
	CookTime     int     `db:"tempo_de_preparo"`
}
