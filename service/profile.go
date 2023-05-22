package service

import (
	"context"
	"github.com/zhayt/student-service/model"
	"github.com/zhayt/student-service/storage"
	"go.uber.org/zap"
)

type IStudentProfileService interface {
	CreateOrUpdateStudentInfo(ctx context.Context, studentInfo model.StudentPersonalInfo) error
	GetStudentAllProfileData(ctx context.Context, studentName string) (model.StudentProfileDTO, error)
}

type StudentProfileService struct {
	storage *storage.Storage
	l       *zap.Logger
}

func NewStudentProfileService(storage *storage.Storage, l *zap.Logger) *StudentProfileService {
	return &StudentProfileService{storage: storage, l: l}
}

func (s *StudentProfileService) GetStudentAllProfileData(ctx context.Context, studentName string) (model.StudentProfileDTO, error) {

	return s.storage.Profile.GetStudentProfileData(ctx, studentName)
}

func (s *StudentProfileService) CreateOrUpdateStudentInfo(ctx context.Context, studentInfo model.StudentPersonalInfo) error {
	// validate data

	return s.storage.Profile.CreateOrUpdateStudentPersonalInfo(ctx, studentInfo)
}
