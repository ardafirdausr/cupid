package entity

import (
	"crypto"
	"time"
)

type UserGender string

const (
	UserGenderMale   UserGender = "male"
	UserGenderFemale UserGender = "female"
)

type User struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Name      string     `json:"name"`
	Bio       string     `json:"bio"`
	Gender    UserGender `json:"gender"`
	BirthDate time.Time  `json:"birth_date"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
}

func (user *User) SetPassword(password string) {
	cryptoAlgo := crypto.SHA256.New()
	cryptoAlgo.Write([]byte(password))
	user.Password = string(cryptoAlgo.Sum(nil))
}
