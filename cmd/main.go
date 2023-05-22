package main

import (
	"fmt"
	"github.com/zhayt/student-service/config"
	"github.com/zhayt/student-service/logger"
	"github.com/zhayt/student-service/service"
	"github.com/zhayt/student-service/storage"
	internal "github.com/zhayt/student-service/transport/http"
	"github.com/zhayt/student-service/transport/http/handler"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

}

func run() error {
	var once sync.Once
	once.Do(config.PrepareENV)

	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	l, err := logger.NewLogger(cfg)
	if err != nil {
		return err
	}

	defer func(l *zap.Logger) {
		err := l.Sync()
		if err != nil {
			log.Fatalln(err)
		}
	}(l)

	repo, err := storage.NewStorage(cfg, l)
	if err != nil {
		return err
	}

	servi := service.NewService(repo, l)

	handle := handler.NewHandler(servi, l)

	httpServer := internal.NewServer(cfg, handle)

	l.Info("Start app", zap.String("port", cfg.AppPort))
	httpServer.Start()

	osSignalCh := make(chan os.Signal, 1)
	signal.Notify(osSignalCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-osSignalCh:
		l.Info("signal accepted: ", zap.String("signal", s.String()))
	case err = <-httpServer.Notify:
		l.Info("server closing", zap.Error(err))
	}

	if err = httpServer.Shutdown(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("error while shutting down server: %s", err)
	}

	return nil
}
