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
	s.e.GET("/player/:id", s.getVote)
	s.e.POST("/player/:id", s.vote)
	admin := s.e.Group("/admin")
	admin.Use(s.adminMiddleware)
	admin.GET("/referee", s.listReferees)
	admin.GET("/referee/:id/", s.listRefereeVotes)
	admin.GET("/switch/:id", s.playerSwitch)
}

func (s *server) checkLogin(c echo.Context) error {
	return c.NoContent(204)
}

func (s *server) getRefereeInfo(c echo.Context) error {
	cc := c.(*CustomContext)
	cc.referee.Key = ""
	return cc.JSON(200, cc.referee)
}

func (s *server) listPlayers(c echo.Context) error {
	cc := c.(*CustomContext)
	Players := make([]Player, 0)
	if result := cc.db.Model(&Player{}).Preload("Votes").Find(&Players); result.Error != nil {
		log.Printf("db error: %s", result.Error)
		return result.Error
	} else {
		if !cc.referee.Admin {
			for k := range Players {
				Players[k].Votes = nil
			}
		} else {
			for k, p := range Players {
				Players[k].Votes = p.GetScore()
			}
		}
		return cc.JSON(200, Players)
	}
}

func (s *server) getVote(c echo.Context) error {
	cc := c.(*CustomContext)
	vote := &Vote{
		VoteBy: int(cc.referee.ID),
	}
	if id, err := strconv.Atoi(cc.Param("id")); err != nil {
		return echo.ErrForbidden
	} else {
		vote.PlayerID = uint(id)
	}
	cc.db.Where(vote).First(vote)
	if vote.ID != 0 {
		return cc.JSON(200, vote)
	} else {
		return echo.ErrNotFound
	}
}

func (s *server) vote(c echo.Context) error {
	cc := c.(*CustomContext)
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
	if !player.Enable {
		return echo.ErrForbidden
	}
	vote := &Vote{
		PlayerID: player.ID,
		VoteBy:   int(cc.referee.ID),
	}
	cc.db.Where(vote).First(vote)
	score := Score{}
	score.loadFromForm(cc)
	vote.Scores = score
	if cc.referee.Main {
		vote.IsMain = true
	}
	if result := cc.db.Save(vote); result.Error != nil {
		log.Printf("db error: %s", result.Error)
		return echo.ErrBadGateway
	} else {
		return cc.JSON(200, vote)
	}
}

func (s *server) playerSwitch(c echo.Context) error {
	cc := c.(*CustomContext)
	player := &Player{}
	if id, err := strconv.Atoi(cc.Param("id")); err != nil {
		return echo.ErrBadRequest
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

func (s *server) listReferees(c echo.Context) error {
	cc := c.(*CustomContext)
	referees := make([]Referee, 0)
	cc.db.Find(&referees)
	return cc.JSON(200, referees)
}

func (s *server) listRefereeVotes(c echo.Context) error {
	cc := c.(*CustomContext)
	votes := make([]Vote, 0)
	if id, err := strconv.Atoi(cc.Param("id")); err != nil {
		return echo.ErrBadRequest
	} else {
		cc.db.Where("vote_by=?", id).Find(&votes)
		return cc.JSON(200, votes)
	}
}
