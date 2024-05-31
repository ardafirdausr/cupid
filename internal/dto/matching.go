package dto

import "com.ardafirdausr.cupid/internal/entity"

type MatchingRecommendationsFilter struct {
	UserID string `json:"userID"`
	Age    int    `json:"age"`
	Limit  int    `json:"limit"`
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
