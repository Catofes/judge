package main

import (
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *server) bindHandler() {
	s.e.HEAD("/", s.checkLogin)
	s.e.GET("/", s.getRefereeInfo)
	s.e.GET("/player", s.listPlayers)
	s.e.POST("/player/:id", s.vote)
	s.e.GET("/admin/switch/:id", s.playerSwitch)
}

func (s *server) checkLogin(c echo.Context) error {
	cc := c.(CustomContext)
	return c.JSON(200, cc.referee)
}

func (s *server) getRefereeInfo(c echo.Context) error {
	cc := c.(CustomContext)
	cc.referee.Key = ""
	return cc.JSON(200, cc.referee)
}

func (s *server) listPlayers(c echo.Context) error {
	cc := c.(CustomContext)
	Players := make([]Player, 0)
	if result := cc.db.Find(&Players); result.Error != nil {
		log.Printf("db error: %s", result.Error)
		return result.Error
	} else {
		if cc.referee.Admin == false {
			for _, p := range Players {
				p.Votes = nil
			}
		} else {
			for _, p := range Players {
				tmp := make([]Vote, 0)
				for _, v := range p.Votes {
					if v.IsMain {
						tmp = append(tmp, v)
					}
				}
				p.Votes = tmp
			}
		}
		return cc.JSON(200, Players)
	}
}

func (s *server) vote(c echo.Context) error {
	cc := c.(CustomContext)
	player := &Player{}
	if id, err := strconv.Atoi(cc.Param("id")); err != nil {
		return echo.ErrForbidden
	} else {
		player.ID = uint(id)
	}
	cc.db.First(player)
	if player.Name == "" {
		return echo.ErrNotFound
	}
	score := Score{}
	score.loadFromForm(cc)
	vote := Vote{
		VoteBy: cc.referee.ID,
		Scores: score,
	}
	if cc.referee.Main {
		vote.IsMain = true
	}
	player.Votes = append(player.Votes, vote)
	if result := cc.db.Save(player); result.Error != nil {
		log.Printf("db error: %s", result.Error)
		return echo.ErrBadGateway
	} else {
		return cc.NoContent(204)
	}
}

func (s *server) playerSwitch(c echo.Context) error {
	cc := c.(CustomContext)
	if cc.referee.Admin == false {
		return echo.ErrForbidden
	}
	player := &Player{}
	if id, err := strconv.Atoi(cc.Param("id")); err != nil {
		return echo.ErrForbidden
	} else {
		player.ID = uint(id)
	}
	cc.db.First(player)
	if player.Name == "" {
		return echo.ErrNotFound
	}
	if cc.QueryParam("start") == "true" {
		player.Enable = true
	} else {
		player.Enable = false
	}
	if result := cc.db.Save(player); result != nil {
		log.Printf("db error: %s", result.Error)
		return echo.ErrBadGateway
	} else {
		return cc.JSON(200, player)
	}
}
