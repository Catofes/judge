package main

import "gorm.io/gorm"

type Referee struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Key   string `gorm:"not null, unique, index"`
	Main  bool
	Admin bool
}
