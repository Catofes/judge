package main

import (
	"log"

	"github.com/tealeg/xlsx/v3"
)

func importPlayer(path, dbpath string) error {
	wb, err := xlsx.OpenFile(path)
	if err != nil {
		log.Fatal(err)
	}
	sh, ok := wb.Sheet["Sheet1"]
	if !ok {
		log.Fatal("Sheet1 does not exist")
	}
	Players := make([]Player, 0)

	for i := 1; i < sh.MaxRow; i++ {
		row, err := sh.Row(i)
		if err != nil {
			log.Fatal(err)
		}
		p := &Player{
			Name: row.GetCell(0).Value,
		}
		Players = append(Players, *p)
	}
	d := (&db{}).init(dbpath).Begin()
	if d.Error != nil {
		log.Fatal(d.Error)
	}
	d.Create(Players)
	if d.Error != nil {
		log.Fatal(d.Error)
	}
	d.Commit()
	return d.Error
}

// func flushAll(dbpath string) error {
// 	d := (&db{}).init(dbpath)
// 	d.Where("1 = 1").Delete(&Player{})
// 	d.Where("1 = 1").Delete(&Referee{})
// 	d.Where("1 = 1").Delete(&Vote{})
// 	return d.Error
// }

func importReferee(path, dbpath string) error {
	wb, err := xlsx.OpenFile(path)
	if err != nil {
		log.Fatal(err)
	}
	sh, ok := wb.Sheet["Sheet1"]
	if !ok {
		log.Fatal("Sheet1 does not exist")
	}
	referee := make([]*Referee, 0)
	for i := 1; i < sh.MaxRow; i++ {
		row, err := sh.Row(i)
		if err != nil {
			log.Fatal(err)
		}
		r := &Referee{
			Name: row.GetCell(0).Value,
			Key:  row.GetCell(1).Value,
		}
		if row.GetCell(2).Value == "是" {
			r.Main = true
		}
		if row.GetCell(3).Value == "是" {
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
