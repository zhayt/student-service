package service

import (
	"github.com/zhayt/student-service/storage"
	"go.uber.org/zap"
)

type Service struct {
	Student IStudentService
	Profile IStudentProfileService
	Image   IImageService
}

func NewService(storage *storage.Storage, l *zap.Logger) *Service {
	serv := &Service{
		Student: NewStudentService(storage, l),
		Profile: NewStudentProfileService(storage, l),
		Image:   NewImageService(storage, l),
	}

	return serv
}
