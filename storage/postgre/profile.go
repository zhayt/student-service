package postgre

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/zhayt/student-service/model"
	"go.uber.org/zap"
)

type StudentProfileStorage struct {
	db *sqlx.DB
	l  *zap.Logger
}

func NewStudentProfileStorage(db *sqlx.DB, l *zap.Logger) *StudentProfileStorage {
	return &StudentProfileStorage{db: db, l: l}
}

func (r *StudentProfileStorage) GetStudentProfileData(ctx context.Context, studentName string) (model.StudentProfileDTO, error) {
	qr := `SELECT sp.full_name, i.image_url, sp.about_student, g.name, sp.country, sp.region, sp.city, sp.birthday_date,
			sp.phone_number
			FROM 
             student s
             INNER JOIN student_personal_info sp ON s.id = sp.student_id
             INNER JOIN image i USING(student_id)
             INNER JOIN gender g ON g.id = sp.gender_id
             WHERE s.name = $1`

	var studentProfile model.StudentProfileDTO

	if err := r.db.GetContext(ctx, &studentProfile, qr, studentName); err != nil {
		return model.StudentProfileDTO{}, fmt.Errorf("couldn't get student profile data: %w", err)
	}

	return studentProfile, nil
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
		studentInfo.Region, studentInfo.City, studentInfo.BirthdayDate, studentInfo.PhoneNumber); err != nil {
		return fmt.Errorf("couldn't create or update student personal info: %w", err)
	}

	return nil
}
