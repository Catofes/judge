package main

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Score struct {
	First  int
	Second int
	Third  int
	Fourth int
	Fifth  int
}

func (s *Score) loadFromForm(cc CustomContext) {
	s.First, _ = strconv.Atoi(cc.FormValue("First"))
	s.Second, _ = strconv.Atoi(cc.FormValue("Second"))
	s.Third, _ = strconv.Atoi(cc.FormValue("Third"))
	s.Fourth, _ = strconv.Atoi(cc.FormValue("Fourth"))
	s.Fifth, _ = strconv.Atoi(cc.FormValue("Fifth"))
}

type Vote struct {
	gorm.Model
	Scores    Score `gorm:"embedded"`
	PlayerID  uint
	VoteBy    uint
	IsMain    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
