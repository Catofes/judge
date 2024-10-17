package main

import (
	"time"

	"gorm.io/gorm"
)

type Score struct {
	First  int
	Second int
	Third  int
}

type Vote struct {
	gorm.Model
	Scores    Score `gorm:"embedded"`
	PlayerID  int
	VoteBy    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
