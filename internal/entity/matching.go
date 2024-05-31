package entity

import "time"

type MatchingStatus int

const (
	MatchingStatusPending  MatchingStatus = 0
	MatchingStatusAccepted MatchingStatus = 1
	MatchingStatusRejected MatchingStatus = 2
)

func (m MatchingStatus) Valid() bool {
	switch m {
	case MatchingStatusPending, MatchingStatusAccepted, MatchingStatusRejected:
		return true
	}

	return false
}

type Matching struct {
	ID          string         `json:"id" bson:"_id"`
	User1ID     string         `json:"user1" bson:"user1ID"`
	User2ID     string         `json:"user2" bson:"user2ID"`
	Status      MatchingStatus `json:"status" bson:"status"`
	MatchedTime time.Time      `json:"matched_time" bson:"matchedTime"`
}
