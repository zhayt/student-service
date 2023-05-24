package service

import (
	"context"
	"github.com/zhayt/student-service/model"
	"github.com/zhayt/student-service/storage"
	"go.uber.org/zap"
	"sync"
)

type IStudentProfileService interface {
	CreateOrUpdateStudentInfo(ctx context.Context, studentInfo model.StudentPersonalInfo) error
	GetStudentAllProfileData(ctx context.Context, studentName string) *model.StudentProfileDTO
}

type StudentProfileService struct {
	storage *storage.Storage
	l       *zap.Logger
}

func NewStudentProfileService(storage *storage.Storage, l *zap.Logger) *StudentProfileService {
	return &StudentProfileService{storage: storage, l: l}
}

func (s *StudentProfileService) GetStudentAllProfileData(ctx context.Context, studentName string) *model.StudentProfileDTO {

	profileDate := &model.StudentProfileDTO{}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		personalInfo, err := s.storage.Profile.GetStudentPersonalInfoByName(ctx, studentName)
		if err != nil {
			s.l.Error("GetStudentPersonalInfoByName error", zap.Error(err))
			profileDate.PersonalInfoType = false
		} else {
			profileDate.PersonalInfoType = true
			profileDate.PersonalInfo = personalInfo
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		image, err := s.storage.Image.GetImageByStudentName(ctx, studentName)
		if err != nil {
			s.l.Error("GetImageByStudentName error", zap.Error(err))
			profileDate.ImageType = false
		} else {
			profileDate.ImageType = true
			profileDate.Image = image
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		gender, err := s.storage.Gender.GetAllGenders(ctx)
		if err != nil {
			s.l.Error("GetAllGenders error", zap.Error(err))
		}

		profileDate.Gender = gender
		wg.Done()
	}()

	wg.Wait()

	return profileDate
}

func (s *StudentProfileService) CreateOrUpdateStudentInfo(ctx context.Context, studentInfo model.StudentPersonalInfo) error {
	// validate data

	return s.storage.Profile.CreateOrUpdateStudentPersonalInfo(ctx, studentInfo)
}
