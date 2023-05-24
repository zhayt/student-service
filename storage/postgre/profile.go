package postgre

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/zhayt/student-service/model"
	"go.uber.org/zap"
	"time"
)

type StudentProfileStorage struct {
	db *sqlx.DB
	l  *zap.Logger
}

func NewStudentProfileStorage(db *sqlx.DB, l *zap.Logger) *StudentProfileStorage {
	return &StudentProfileStorage{db: db, l: l}
}

func (r *StudentProfileStorage) GetStudentPersonalInfoByName(ctx context.Context, studentName string) (model.StudentPersonalInfoDTO, error) {
	qr := `SELECT full_name, about_student, g.name, country, region, city, birthday_date, phone_number
			FROM student
			INNER JOIN student_personal_info spi ON student.id = spi.student_id
			INNER JOIN gender g ON spi.gender_id = g.id
			WHERE student.name = $1
			`

	var personalInfo model.StudentPersonalInfoDTO

	if err := r.db.GetContext(ctx, &personalInfo, qr, studentName); err != nil {
		return model.StudentPersonalInfoDTO{}, fmt.Errorf("couldn't get student personal info: %w", err)
	}

	return personalInfo, nil
}

func (r *StudentProfileStorage) CreateOrUpdateStudentPersonalInfo(ctx context.Context, studentInfo model.StudentPersonalInfo) error {
	qr := `INSERT INTO student_personal_info 
    (student_id, gender_id, full_name, about_student, country, region, city, birthday_date, phone_number)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    ON CONFLICT (student_id)
    DO UPDATE 
        SET gender_id = EXCLUDED.gender_id, 
    		full_name = EXCLUDED.full_name, 
    		about_student = EXCLUDED.about_student, 
    		country = EXCLUDED.country, 
    		region = EXCLUDED.region, 
    		city = EXCLUDED.city, 
    		birthday_date = EXCLUDED.birthday_date, 
    		phone_number = EXCLUDED.phone_number;
        `

	if _, err := r.db.ExecContext(ctx, qr,
		studentInfo.StudentID, studentInfo.GenderID, studentInfo.FullName, studentInfo.AboutStudent, studentInfo.Country,
		studentInfo.Region, studentInfo.City, time.Time(studentInfo.BirthdayDate), studentInfo.PhoneNumber); err != nil {
		return fmt.Errorf("couldn't create or update student personal info: %w", err)
	}

	return nil
}
