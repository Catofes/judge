package main

import "github.com/labstack/echo/v4"

func (s *server) bindHandler() {
	s.e.HEAD("/", s.checkLogin)
	s.e.GET("/", s.getRefereeInfo)
}

func (s *server) checkLogin(c echo.Context) error {
	return c.NoContent(204)
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
		return result.Error
	} else {
		return cc.JSON(200, Players)
	}
}
