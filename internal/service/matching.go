package service

import (
	"context"
	"time"

	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/entity"
	"com.ardafirdausr.cupid/internal/entity/errs"
	customJWT "com.ardafirdausr.cupid/internal/pkg/jwt"
	"github.com/pkg/errors"
)

type MatchingService struct {
	matchingRepo internal.MatchingRepositorier
	userRepo     internal.UserRepositorier
}

func NewMatchingService(matchingRepo internal.MatchingRepositorier, userRepo internal.UserRepositorier) *MatchingService {
	return &MatchingService{matchingRepo: matchingRepo, userRepo: userRepo}
}

func (svc *MatchingService) GetMatchingRecommendations(ctx context.Context, filter dto.MatchingRecommendationsFilter) ([]entity.User, error) {
	// TODO: add recommendation by user's preference
	users, err := svc.matchingRepo.GetMatchingRecommendations(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get matching recommendations")
	}

	return users, nil
}

func (svc *MatchingService) MatchMaking(ctx context.Context, param dto.CreateMatchingParam) (*entity.Matching, error) {
	user, err := customJWT.GetUserFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user from context")
	}

	if param.UserID == user.ID {
		return nil, errs.NewErrUnprocessable("cannot match with yourself 🥰")
	}

	if isAble, err := svc.isUserAbleToMakeMatchMaking(ctx, user); err != nil {
		return nil, errors.Wrap(err, "failed to check user ability")
	} else if !isAble {
		return nil, errs.NewErrUnprocessable("user is not able to make matchmaking")
	}

	// Get user
	if _, err := svc.userRepo.GetUserByID(ctx, param.UserID); err != nil {
		return nil, errors.Wrap(err, "failed to get target user")
	}

	// Get matching
	matching, err := svc.matchingRepo.GetMatchingByUser(ctx, user.ID, param.UserID)
	if err != nil && !errs.IsEqualType(err, errs.ErrNotFound) {
		return nil, errors.Wrap(err, "failed to get matching")
	}

	now := time.Now()
	nextStatus := entity.MatchingStatus(param.Status)

	// matchmaking (second move)
	if matching != nil {
		if matching.User1ID == user.ID && matching.User1SwapAt != nil {
			return nil, errs.NewErrUnprocessable("user already swiped")
		}

		switch matching.Status {
		case entity.MatchingStatusRejected, entity.MatchingStatusMatched:
			return nil, errs.NewErrInvalidData("matching already processed")
		case entity.MatchingStatusAccepted:
			if entity.MatchingStatus(param.Status) == entity.MatchingStatusAccepted {
				nextStatus = entity.MatchingStatusMatched
			} else if entity.MatchingStatus(param.Status) == entity.MatchingStatusRejected {
				nextStatus = entity.MatchingStatusRejected
			}
		}

		// update matching
		matching.User2SwapAt = &now
		matching.Status = nextStatus
		if nextStatus == entity.MatchingStatusMatched {
			matching.MatchedAt = &now
		}

		if err := svc.matchingRepo.UpdateMatchingByID(ctx, matching.ID, matching); err != nil {
			return nil, errors.Wrap(err, "failed to create matching")
		}

		return matching, nil

	}

	// new matchmaking (first move)
	matching = &entity.Matching{
		User1ID:     user.ID,
		User1SwapAt: &now,
		User2ID:     param.UserID,
		User2SwapAt: nil,
		Status:      nextStatus,
	}

	if err := svc.matchingRepo.CreateMatching(ctx, matching); err != nil {
		return nil, errors.Wrap(err, "failed to create matching")
	}

	return matching, nil
}

func (svc *MatchingService) isUserAbleToMakeMatchMaking(ctx context.Context, user *entity.User) (bool, error) {
	acceptedCount, err := svc.matchingRepo.GetUserMatchingCount(ctx, user.ID, time.Now())
	if err != nil {
		return false, errors.Wrap(err, "failed to get user accepted count")
	}

	if acceptedCount >= 10 {
		return false, errs.NewErrUnprocessable("user has reached maximum accepted count")
	}

	return true, nil
}
