package dto

import (
	"time"

	"com.ardafirdausr.cupid/internal/entity"
	"com.ardafirdausr.cupid/internal/entity/errs"
)

type LoginUserParam struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterUserParam struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6,max=50"`
	Name      string `json:"name" validate:"required,min=3,max=50"`
	Bio       string `json:"bio" validate:"omitempty,min=10,max=255"`
	BirthDate string `json:"birth_date" validate:"required,datetime=2006-01-02"`
	Gender    string `json:"gender" validate:"required,oneof=male female"`
}

func (p *RegisterUserParam) ToUser(user *entity.User) (err error) {
	user.Email = p.Email
	user.Name = p.Name
	user.Bio = p.Bio
	user.Gender = entity.UserGender(p.Gender)
	if user.BirthDate, err = time.Parse(time.DateOnly, p.BirthDate); err != nil {
		return errs.NewErrInvalidData("birth_date", "invalid date format")
	}

	user.SetPassword(p.Password)
	return nil
}
