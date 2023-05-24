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
	GetStudentPersonalInfoByName(ctx context.Context, studentName string) (model.StudentPersonalInfoDTO, error)
}

type IImageStorage interface {
	GetImageByStudentName(ctx context.Context, studentName string) (model.Image, error)
	CreateOrUpdateImage(ctx context.Context, image model.Image) error
	DeleteImageByStudentID(ctx context.Context, studentID int) error
}

type IGenderStorage interface {
	GetAllGenders(ctx context.Context) ([]*model.Gender, error)
}

type Storage struct {
	Student IStudentStorage
	Profile IStudentProfileStorage
	Image   IImageStorage
	Gender  IGenderStorage
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
		Gender:  postgre.NewGenderStorage(db, l),
	}

	return store, nil
}
