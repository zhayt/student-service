package postgre

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/zhayt/student-service/model"
	"go.uber.org/zap"
)

type ImageStorage struct {
	db *sqlx.DB
	l  *zap.Logger
}

func NewImageStorage(db *sqlx.DB, l *zap.Logger) *ImageStorage {
	return &ImageStorage{db: db, l: l}
}

func (r *ImageStorage) CreateImage(ctx context.Context, image model.Image) (int, error) {
	qr := `INSERT INTO image(student_id, image_url) VALUES ($1, $2) RETURNING id`

	var imageID int

	if err := r.db.GetContext(ctx, &imageID, qr, image.StudentID, image.ImageURL); err != nil {
		return 0, fmt.Errorf("couldn't crreate image: %w", err)
	}

	return imageID, nil
}

func (r *ImageStorage) UpdateImage(ctx context.Context, image model.Image) error {
	qr := `UPDATE image SET image_url = $1 WHERE student_id = $2`

	if _, err := r.db.ExecContext(ctx, qr, image.ImageURL, image.StudentID); err != nil {
		return fmt.Errorf("couldn't update image: %w", err)
	}

	return nil
}

func (r *ImageStorage) DeleteImageByStudentID(ctx context.Context, studentID int) error {
	qr := `DELETE FROM image WHERE student_id = $1`

	if _, err := r.db.Exec(qr, studentID); err != nil {
		return fmt.Errorf("couldn't delete image: %w", err)
	}

	return nil
}
