package entity

type Weekdays uint8

const (
	Monday    Weekdays = 1
	Tuesday   Weekdays = 2
	Wednesday Weekdays = 4
	Thursday  Weekdays = 8
	Friday    Weekdays = 16
	Saturday  Weekdays = 32
	Sunday    Weekdays = 64
)

type Restaurant struct {
	Id          int       `db:"id"`
	OpenHour    string    `db:"horario_abertura"`
	CloseHour   string 	  `db:"horario_fechamento"`
	CityID      *int      `db:"id_cidade"`
	OpenDays    Weekdays  `db:"dias_funcionamento"`
	Name        string    `db:"nome"`
	Description *string   `db:"descricao"`
	PhoneNumber string    `db:"telefone"`
	Address     string    `db:"endereco"`
}

func (r *Restaurant) IsEmpty() bool {
	return r == nil || r.Id == 0
}
