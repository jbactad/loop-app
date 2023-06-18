package infrastructure

import "github.com/google/uuid"

type UUIDGeneratorImpl struct{}

func NewUUIDGenerator() *UUIDGeneratorImpl {
	return &UUIDGeneratorImpl{}
}

func (ug UUIDGeneratorImpl) Generate() string {
	return uuid.New().String()
}
