package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golangProject.com/localization/repository"
	"log"
)

const dbName = "test"

type Store struct {
	logger log.Logger
	client *mongo.Client

	LocaleRepository repository.LocaleRepository
}

func NewStore(client *mongo.Client) (*Store, error) {
	store := &Store{
		client: client,
	}

	err := store.migrate()
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *Store) Locale() repository.LocaleRepository {
	if s.LocaleRepository != nil {
		return s.LocaleRepository
	}

	s.LocaleRepository = &LocaleRepo{
		store: s,
	}

	return s.LocaleRepository
}
