package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type db struct {
	*gorm.DB
}

func (s *db) init(path string) *db {
	var err error
	s.DB, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	s.migrate()
	return s
}

func (s *db) migrate() {
	if err := s.AutoMigrate(Vote{}, Player{}, Referee{}); err != nil {
		log.Fatal(err)
	}
}
