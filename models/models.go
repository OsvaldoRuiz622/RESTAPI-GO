package models

import "gorm.io/gorm"

type Sabores struct {
	gorm.Model

	Sabor string `gorm:"text;not null;default:null" json:"Sabor"`
}
