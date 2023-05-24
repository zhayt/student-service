package model

import (
	"encoding/json"
	"strings"
	"time"
)

type Student struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StudentDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Gender struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type StudentPersonalInfo struct {
	ID           int           `json:"id" db:"id"`
	StudentID    int           `json:"studentID" db:"student_id"`
	GenderID     int           `json:"genderID" db:"gender_id"`
	FullName     string        `json:"fullName" db:"full_name"`
	AboutStudent string        `json:"aboutStudent" db:"about_student"`
	Country      string        `json:"country" db:"country"`
	Region       string        `json:"region" db:"region"`
	City         string        `json:"city" db:"city"`
	BirthdayDate JsonBirthDate `json:"birthdayDate" db:"birthday_date"`
	PhoneNumber  string        `json:"phoneNumber" db:"phone_number"`
}

type Image struct {
	ID        int    `json:"-"`
	StudentID int    `json:"studentID" db:"student_id"`
	ImageURL  string `json:"imageURL" db:"image_url"`
}

type StudentPersonalInfoDTO struct {
	FullName     string    `json:"fullName" db:"full_name"`
	AboutStudent string    `json:"aboutStudent" db:"about_student"`
	GenderName   string    `json:"genderName" db:"name"`
	Country      string    `json:"country" db:"country"`
	Region       string    `json:"region" db:"region"`
	City         string    `json:"city" db:"city"`
	BirthdayDate time.Time `json:"birthdayDate" db:"birthday_date"`
	PhoneNumber  string    `json:"phoneNumber" db:"phone_number"`
}

type StudentProfileDTO struct {
	PersonalInfoType bool                   `json:"personalInfoType"`
	PersonalInfo     StudentPersonalInfoDTO `json:"personalInfo"`
	ImageType        bool                   `json:"imageType"`
	Image            Image                  `json:"image"`
}

type JsonBirthDate time.Time

// Implement Marshaler and Unmarshaler interface
func (j *JsonBirthDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonBirthDate(t)
	return nil
}

func (j JsonBirthDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

// Maybe a Format function for printing your date
func (j JsonBirthDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
