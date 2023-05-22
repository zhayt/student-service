package postgre

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/zhayt/student-service/config"
	"time"
)

func Dial(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect(cfg.DBDriver, cfg.DBConURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't get pool connection: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("couldn't connect to db: %w", err)
	}

	return db, nil
}
