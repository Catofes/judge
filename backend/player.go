package main

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name   string
	Enable bool
	Votes  []Vote
}
