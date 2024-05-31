package mongo

import "time"

type Config struct {
	URI             string
	DB              string
	MinPool         uint64
	MaxPool         uint64
	MaxIdleTimePool time.Duration
}
