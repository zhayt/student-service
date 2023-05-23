package postgre

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/zhayt/student-service/model"
	"go.uber.org/zap"
)

type StudentStorage struct {
	db *sqlx.DB
	l  *zap.Logger
}

func NewStudentStorage(db *sqlx.DB, l *zap.Logger) *StudentStorage {
	return &StudentStorage{db: db, l: l}
}

func (r *StudentStorage) CreateStudent(ctx context.Context, student model.Student) (int, error) {
	qr := `INSERT INTO student(name, email, password) VALUES ($1, $2, $3) RETURNING id`

	var studentID int

	if err := r.db.GetContext(ctx, &studentID, qr, student.Name, student.Email, student.Password); err != nil {
		return 0, fmt.Errorf("couldn't create student: %w", err)
	}

	return studentID, nil
}

func (r *StudentStorage) GetStudentByEmail(ctx context.Context, studentEmail string) (model.Student, error) {
	qr := `SELECT * FROM student WHERE email = $1`

	r.l.Info("Try get data by email", zap.String("email", studentEmail))

	var student model.Student

	if err := r.db.GetContext(ctx, &student, qr, studentEmail); err != nil {
		return model.Student{}, fmt.Errorf("couldn't get student by email: %w", err)
	}

	return student, nil
}
