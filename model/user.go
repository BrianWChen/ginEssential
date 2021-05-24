package model

import "gorm.io/gorm"

type Group string

const (
	Admin   Group = "admin"
	Manager Group = "manager"
	Enduser Group = "user"
	Guest   Group = "group"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(100);not null;unique"`
	Password  string `gorm:"size:255;not null"`
	Group     Group  `gorm:"type:enum('admin', 'manager', 'user', 'guest');default:'guest'"`
}
