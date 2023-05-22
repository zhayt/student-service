package storage

import (
	"context"
	"github.com/zhayt/student-service/config"
	"github.com/zhayt/student-service/model"
	"github.com/zhayt/student-service/storage/postgre"
	"go.uber.org/zap"
)

type IStudentStorage interface {
	CreateStudent(ctx context.Context, student model.Student) (int, error)
	GetStudentByEmail(ctx context.Context, studentEmail string) (model.Student, error)
}

type IStudentProfileStorage interface {
	CreateOrUpdateStudentPersonalInfo(ctx context.Context, studentInfo model.StudentPersonalInfo) error
	GetStudentProfileData(ctx context.Context, studentName string) (model.StudentProfileDTO, error)
}

type IImageStorage interface {
	CreateImage(ctx context.Context, image model.Image) (int, error)
	UpdateImage(ctx context.Context, image model.Image) error
	DeleteImageByStudentID(ctx context.Context, studentID int) error
}

type Storage struct {
	Student IStudentStorage
	Profile IStudentProfileStorage
	Image   IImageStorage
}

func NewStorage(cfg *config.Config, l *zap.Logger) (*Storage, error) {
	db, err := postgre.Dial(cfg)
	if err != nil {
		return nil, err
	}

	store := &Storage{
		Student: postgre.NewStudentStorage(db, l),
		Profile: postgre.NewStudentProfileStorage(db, l),
		Image:   postgre.NewImageStorage(db, l),
	}

	return store, nil
}
