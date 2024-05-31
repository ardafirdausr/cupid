package entity

import (
	"crypto"
	"encoding/base64"
	"time"
)

type UserGender string

const (
	UserGenderMale   UserGender = "male"
	UserGenderFemale UserGender = "female"
)

type User struct {
	ID        string     `json:"id" bson:"_id"`
	Email     string     `json:"email" bson:"email"`
	Password  string     `json:"-" bson:"password" `
	Name      string     `json:"name" bson:"name" `
	Bio       string     `json:"bio" bson:"bio" `
	Gender    UserGender `json:"gender" bson:"gender"`
	BirthDate time.Time  `json:"birth_date" bson:"birthDate"`
}

func (user *User) Age() int {
	now := time.Now()
	age := now.Year() - user.BirthDate.Year()
	if now.YearDay() < user.BirthDate.YearDay() {
		age--
	}

	return age
}

func (user *User) SetPassword(password string) {
	user.Password = user.hashPassword(password)
}

func (user *User) ComparePassword(password string) bool {
	return user.Password == user.hashPassword(password)
}

func (user *User) hashPassword(password string) string {
	cryptoAlgo := crypto.SHA256.New()
	cryptoAlgo.Write([]byte(password))
	return base64.StdEncoding.EncodeToString(cryptoAlgo.Sum(nil))
}
