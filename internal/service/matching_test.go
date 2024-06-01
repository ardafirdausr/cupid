package service

import (
	"context"
	"testing"

	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/entity"
	"com.ardafirdausr.cupid/internal/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetMatchingRecommendations(t *testing.T) {
	t.Run("failed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		filter := dto.MatchingRecommendationsFilter{
			UserID: "user-id",
		}

		mockMatchingRepo := mock.NewMockMatchingRepositorier(ctrl)
		mockMatchingRepo.EXPECT().GetMatchingRecommendations(gomock.Any(), filter).Return(nil, assert.AnError)

		service := NewMatchingService(mockMatchingRepo, nil)
		users, err := service.GetMatchingRecommendations(context.Background(), filter)

		assert.Error(t, err)
		assert.Nil(t, users)
	})

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		filter := dto.MatchingRecommendationsFilter{
			UserID: "user-id",
		}

		expectedUsers := []entity.User{
			{
				ID: "user-id-1",
			},
			{
				ID: "user-id-2",
			},
		}

		mockMatchingRepo := mock.NewMockMatchingRepositorier(ctrl)
		mockMatchingRepo.EXPECT().GetMatchingRecommendations(gomock.Any(), filter).Return(expectedUsers, nil)

		service := NewMatchingService(mockMatchingRepo, nil)
		users, err := service.GetMatchingRecommendations(context.Background(), filter)

		assert.NoError(t, err)
		assert.Equal(t, expectedUsers, users)
	})
}
