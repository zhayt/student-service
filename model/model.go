package model

import "time"

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
	ID           int       `json:"id" db:"id"`
	StudentID    int       `json:"studentID" db:"student_id"`
	GenderID     int       `json:"genderID" db:"gender_id"`
	FullName     string    `json:"fullName" db:"full_name"`
	AboutStudent string    `json:"aboutStudent" db:"about_student"`
	Country      string    `json:"country" db:"country"`
	Region       string    `json:"region" db:"region"`
	City         string    `json:"city" db:"city"`
	BirthdayDate time.Time `json:"birthdayDate" db:"birthday_date"`
	PhoneNumber  string    `json:"phoneNumber" db:"phone_number"`
}

type Image struct {
	ID        int    `json:"-"`
	StudentID int    `json:"studentID" db:"student_id"`
	ImageURL  string `json:"imageURL" db:"image_url"`
}

type StudentProfileDTO struct {
	FullName     string    `json:"fullName" db:"full_name"`
	ImageURL     string    `json:"imageURL" db:"image_url"`
	AboutStudent string    `json:"aboutStudent" db:"about_student"`
	GenderName   string    `json:"genderName" db:"name"`
	Country      string    `json:"country" db:"country"`
	Region       string    `json:"region" db:"region"`
	City         string    `json:"city" db:"city"`
	BirthdayDate time.Time `json:"birthdayDate" db:"birthday_date"`
	PhoneNumber  string    `json:"phoneNumber" db:"phone_number"`
}
