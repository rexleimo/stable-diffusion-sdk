package models

import "gorm.io/gorm"

type UserCheckInToday struct {
	gorm.Model
	UserID string
}
