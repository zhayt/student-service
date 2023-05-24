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

func (r *ImageStorage) GetImageByStudentName(ctx context.Context, studentName string) (model.Image, error) {
	qr := `SELECT image.id, image.student_id, image.image_url
			FROM image
			INNER JOIN student ON student.id = image.student_id
			WHERE student.name = $1`

	var image model.Image

	if err := r.db.GetContext(ctx, &image, qr, studentName); err != nil {
		return model.Image{}, fmt.Errorf("couldn't get image: %w", err)
	}

	return image, nil
}

func (r *ImageStorage) CreateOrUpdateImage(ctx context.Context, image model.Image) error {
	qr := `INSERT INTO image (student_id, image_url) 
			VALUES ($1, $2)
			ON CONFLICT (student_id)
			DO UPDATE 
			    SET image_url = EXCLUDED.image_url`

	r.l.Info("Try to create or update image", zap.Any("imageObj", image))

	if _, err := r.db.ExecContext(ctx, qr, image.StudentID, image.ImageURL); err != nil {
		return fmt.Errorf("couldn't create or update image: %w", err)
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
