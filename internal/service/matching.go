package service

import (
	"context"

	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/entity"
	"github.com/pkg/errors"
)

type MatchingService struct {
	matchingRepo internal.MatchingRepositorier
}

func NewMatchingService(matchingRepo internal.MatchingRepositorier) *MatchingService {
	return &MatchingService{matchingRepo: matchingRepo}
}

func (svc *MatchingService) GetMatchingRecommendations(ctx context.Context, filter dto.MatchingRecommendationsFilter) ([]entity.User, error) {
	// TODO: add recommendation by user's preference
	users, err := svc.matchingRepo.GetMatchingRecommendations(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get matching recommendations")
	}

	return users, nil
}

func (svc *MatchingService) CreateMatching(ctx context.Context, param dto.CreateMatchingParam) (*entity.Matching, error) {
	return nil, nil
}
