package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zhayt/student-service/config"
	"github.com/zhayt/student-service/transport/http/handler"
	"time"
)

const (
	_defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	App             *echo.Echo
	cfg             *config.Config
	handler         *handler.Handler
	Notify          chan error
	shutdownTimeout time.Duration
}

func NewServer(cfg *config.Config, handler *handler.Handler) *Server {
	srv := &Server{
		cfg:             cfg,
		handler:         handler,
		shutdownTimeout: _defaultShutdownTimeout,
		Notify:          make(chan error, 1),
	}

	return srv
}

func (s *Server) BuildingEngine() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	return e
}

func (s *Server) Start() {
	s.App = s.BuildingEngine()
	s.SetUpRoute()
	go func() {
		s.Notify <- s.App.Start(fmt.Sprintf(":%s", s.cfg.AppPort))
		close(s.Notify)
	}()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return s.App.Shutdown(ctx)
}
