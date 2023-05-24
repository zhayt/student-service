package postgre

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/zhayt/student-service/model"
	"go.uber.org/zap"
)

type GenderStorage struct {
	db *sqlx.DB
	l  *zap.Logger
}

func NewGenderStorage(db *sqlx.DB, l *zap.Logger) *GenderStorage {
	return &GenderStorage{db: db, l: l}
}

func (r *GenderStorage) GetAllGenders(ctx context.Context) ([]*model.Gender, error) {
	qr := `SELECT * FROM gender`

	var genders []*model.Gender

	if err := r.db.SelectContext(ctx, &genders, qr); err != nil {
		return []*model.Gender{}, err
	}

	return genders, nil
}
