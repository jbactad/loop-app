package infrastructure

import (
	"time"

	"github.com/google/uuid"
)

type GoogleUUIDGenerator struct{}

func NewGoogleUUIDGenerator() *GoogleUUIDGenerator {
	return &GoogleUUIDGenerator{}
}

func (ug GoogleUUIDGenerator) Generate() string {
	return uuid.New().String()
}

type TimeProvider struct{}

func NewTimeProvider() (provider *TimeProvider, err error) {
	provider = &TimeProvider{}
	return
}

func (tp TimeProvider) Now() time.Time {
	return time.Now()
}
