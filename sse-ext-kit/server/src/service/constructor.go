package service

import (
	"github.com/go-kit/log"
	"server/src/repository"
)

type service struct {
	store  repository.Storage
	logger log.Logger
}

func NewService(
	store repository.Storage,
	logger log.Logger,
) Service {
	return &service{
		store:  store,
		logger: logger,
	}
}
