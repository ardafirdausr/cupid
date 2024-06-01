package dto

import "com.ardafirdausr.cupid/internal/entity"

type MatchingRecommendationsFilter struct {
	UserID string
	Age    int `query:"age"`
	Gender string
	Limit  int `query:"limit"`
}

func (p *MatchingRecommendationsFilter) SetDefault() {
	if p.Limit == 0 {
		p.Limit = 10
	}
}

func (p *MatchingRecommendationsFilter) SetByUser(user *entity.User) {
	p.UserID = user.ID
	switch user.Gender {
	case entity.UserGenderFemale:
		p.Gender = string(entity.UserGenderMale)
	case entity.UserGenderMale:
		p.Gender = string(entity.UserGenderFemale)
	}
}

type CreateMatchingParam struct {
	UserID string `json:"userID" validate:"required"`
	Status int    `json:"status" validate:"required,oneof=1 2"`
}

func (p *CreateMatchingParam) ToMatching(user *entity.User, matching *entity.Matching) {
	matching.User1ID = user.ID
	matching.User2ID = p.UserID
	matching.Status = entity.MatchingStatus(p.Status)
}
