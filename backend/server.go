package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomContext struct {
	echo.Context
	db      *gorm.DB
	referee *Referee
}

type server struct {
	e              *echo.Echo
	d              *db
	Listen         string
	Database       string
	StaticFilePath string
}

func (s *server) init() {
	s.e = echo.New()
	s.d = (&db{}).init(s.Database)
	s.bind()
}

func (s *server) serve() {
	s.e.Logger.Fatal(s.e.Start(s.Listen))
}

func (s *server) bind() {
	s.bindHandler()
}

func (s *server) dbMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{c, nil, nil}
		cc.db = s.d.Begin()
		err := next(cc)
		if err != nil || (cc.Response().Status != 200 && cc.Response().Status != 204) {
			cc.db.Rollback()
		} else {
			cc.db.Commit()
		}
		return err
	}
}

func (s *server) authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*CustomContext)
		cc.referee = &Referee{}
		if key := cc.Request().Header.Get("key"); key == "" {
			return echo.ErrForbidden
		} else {
			cc.referee.Key = key
			if cc.referee.Key == "" {
				return echo.ErrForbidden
			}
			cc.db.Where(cc.referee).First(cc.referee)
			if cc.referee.ID == 0 {
				return echo.ErrNotFound
			}
		}
		return next(cc)
	}
}

func (s *server) adminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*CustomContext)
		if !cc.referee.Admin {
			return echo.ErrForbidden
		}
		return next(cc)
	}
}

func (s *server) nocacheMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	Ext := func(path string) string {
		for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
			if path[i] == '.' {
				return path[i:]
			}
		}
		return ""
	}
	return func(c echo.Context) error {
		if Ext(c.Request().URL.Path); false {
			//if end := Ext(c.Request().URL.Path); end == "js" || end == "html" || end == "css" {
			c.Response().Header().Set("Cache-Control", "max-age=120")
		} else {
			c.Response().Header().Set("Cache-Control", "no-cache")
		}
		return next(c)
	}
}
