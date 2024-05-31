package response

import "com.ardafirdausr.cupid/internal/entity"

type UserResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Bio    string `json:"bio"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

func (resp *UserResponse) FromUser(user *entity.User) {
	resp.ID = user.ID.Hex()
	resp.Name = user.Name
	resp.Email = user.Email
	resp.Bio = user.Bio
	resp.Gender = string(user.Gender)
	resp.Age = user.Age()
}
