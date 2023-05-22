package service

import (
	"context"
	"github.com/zhayt/student-service/model"
	"github.com/zhayt/student-service/storage"
	"go.uber.org/zap"
)

type IImageService interface {
	CreateImage(ctx context.Context, image model.Image) (int, error)
	UpdateImage(ctx context.Context, image model.Image) error
	DeleteImage(ctx context.Context, studentID int) error
}

type ImageService struct {
	storage *storage.Storage
	l       *zap.Logger
}

func NewImageService(storage *storage.Storage, l *zap.Logger) *ImageService {
	return &ImageService{storage: storage, l: l}
}

func (s *ImageService) CreateImage(ctx context.Context, image model.Image) (int, error) {
	return s.storage.Image.CreateImage(ctx, image)
}

func (s *ImageService) UpdateImage(ctx context.Context, image model.Image) error {
	return s.UpdateImage(ctx, image)
}

func (s *ImageService) DeleteImage(ctx context.Context, studentID int) error {
	go func() {
		if err := s.storage.Image.DeleteImageByStudentID(ctx, studentID); err != nil {
			s.l.Error("DeleteImageByStudentID error", zap.Error(err))
		}
	}()

	return nil
}
