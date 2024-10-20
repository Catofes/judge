package main

import (
	"encoding/csv"
	"log"
	"os"
)

func importPlayer(path, dbpath string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	players := make([]Player, 0)
	for _, row := range data[1:] {
		p := Player{
			Name:   row[0],
			Enable: false,
		}
		players = append(players, p)
	}
	d := (&db{}).init(dbpath).Begin()
	if d.Error != nil {
		log.Fatal(d.Error)
	}
	d.Create(players)
	if d.Error != nil {
		log.Fatal(d.Error)
	}
	d.Commit()
	return d.Error
}

func flushAll(dbpath string) error {
	d := (&db{}).init(dbpath)
	d.Where("1 = 1").Delete(&Player{})
	d.Where("1 = 1").Delete(&Referee{})
	d.Where("1 = 1").Delete(&Vote{})
	return d.Error
}

func importReferee(path, dbpath string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	referee := make([]Referee, 0)
	for _, row := range data[1:] {
		r := Referee{
			Name: row[0],
			Key:  row[1],
		}
		if row[2] == "是" {
			r.Main = true
		}
		if row[3] == "是" {
			r.Admin = true
		}
		referee = append(referee, r)
	}
	d := (&db{}).init(dbpath).Begin()
	if d.Error != nil {
		log.Fatal(d.Error)
	}
	d.Create(referee)
	if d.Error != nil {
		log.Fatal(d.Error)
	}
	d.Commit()
	return nil
}
