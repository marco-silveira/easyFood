package models

import (
	"easyfood/pkg/entity"
)

func NewRestaurant(restaurants ...*entity.Restaurant) []*Restaurant {
	result := make([]*Restaurant, 0)
	for _, restaurant := range restaurants {
		result = append(result, &Restaurant{
			ID:          restaurant.Id,
			Name:        restaurant.Name,
			Description: restaurant.Description,
			OpenDays:    GetModelsWeekdays(restaurant.OpenDays),
			PhoneNumber: restaurant.PhoneNumber,
			Address:     restaurant.Address,
			OpenHour:    restaurant.OpenHour,
			CloseHour:   restaurant.CloseHour,
		})
	}

	return result
}

func GetEntityWeekdays(weekdays []Weekdays) entity.Weekdays {
	var result entity.Weekdays
	for _, wd := range weekdays {
		switch wd {
		case WeekdaysMonday:
			result = result | entity.Monday
		case WeekdaysTuesday:
			result = result | entity.Tuesday
		case WeekdaysWednesday:
			result = result | entity.Wednesday
		case WeekdaysThursday:
			result = result | entity.Thursday
		case WeekdaysFriday:
			result = result | entity.Friday
		case WeekdaysSaturday:
			result = result | entity.Saturday
		case WeekdaysSunday:
			result = result | entity.Sunday
		}
	}

	return result
}

func GetModelsWeekdays(weekdays entity.Weekdays) []Weekdays {
	result := make([]Weekdays, 0)
	if weekdays&entity.Monday == entity.Monday {
		result = append(result, WeekdaysMonday)
	}
	if weekdays&entity.Tuesday == entity.Tuesday {
		result = append(result, WeekdaysTuesday)
	}
	if weekdays&entity.Wednesday == entity.Wednesday {
		result = append(result, WeekdaysWednesday)
	}
	if weekdays&entity.Thursday == entity.Thursday {
		result = append(result, WeekdaysThursday)
	}
	if weekdays&entity.Friday == entity.Friday {
		result = append(result, WeekdaysFriday)
	}
	if weekdays&entity.Saturday == entity.Saturday {
		result = append(result, WeekdaysSaturday)
	}
	if weekdays&entity.Sunday == entity.Sunday {
		result = append(result, WeekdaysSunday)
	}

	return result
}
