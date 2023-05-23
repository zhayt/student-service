package service

import (
	"context"
	"fmt"
	"github.com/zhayt/student-service/model"
	"github.com/zhayt/student-service/storage"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type IStudentService interface {
	CreateStudent(ctx context.Context, student model.Student) (int, error)
	GetStudentByEmail(ctx context.Context, studentTDO model.StudentDTO) (model.Student, error)
}

type StudentService struct {
	storage *storage.Storage
	l       *zap.Logger
}

func NewStudentService(storage *storage.Storage, l *zap.Logger) *StudentService {
	return &StudentService{storage: storage, l: l}
}

func (s *StudentService) CreateStudent(ctx context.Context, student model.Student) (int, error) {
	// go validation
	passwordHash, err := generatePasswordHash(student.Password)
	if err != nil {
		return 0, fmt.Errorf("generete password hash error: %w", err)
	}

	student.Password = passwordHash

	return s.storage.Student.CreateStudent(ctx, student)
}

func (s *StudentService) GetStudentByEmail(ctx context.Context, studentDTO model.StudentDTO) (model.Student, error) {
	student, err := s.storage.Student.GetStudentByEmail(ctx, studentDTO.Email)
	if err != nil {
		return model.Student{}, err
	}

	err = compareHashAndPassword(student.Password, studentDTO.Password)
	if err != nil {
		return model.Student{}, fmt.Errorf("compare hash and password error: %w", err)
	}

	return student, nil
}

func generatePasswordHash(passwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("can't generate password hash: %w", err)
	}

	return string(hash), nil
}

func compareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
