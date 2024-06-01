package entity

import "time"

type MatchingStatus int

const (
	MatchingStatusRejected MatchingStatus = 1
	MatchingStatusAccepted MatchingStatus = 2
	MatchingStatusMatched  MatchingStatus = 3
)

func (m MatchingStatus) Valid() bool {
	switch m {
	case MatchingStatusMatched, MatchingStatusAccepted, MatchingStatusRejected:
		return true
	}

	return false
}

type Matching struct {
	ID          string         `json:"id" bson:"_id"`
	User1ID     string         `json:"user1" bson:"user1ID"`
	User1SwapAt *time.Time     `json:"user1_swap_at" bson:"user1SwapAt"`
	User2ID     string         `json:"user2" bson:"user2ID"`
	User2SwapAt *time.Time     `json:"user2_swap_at" bson:"user2SwapAt"`
	Status      MatchingStatus `json:"status" bson:"status"`
	MatchedAt   *time.Time     `json:"matched_time" bson:"matchedTime"`
}
