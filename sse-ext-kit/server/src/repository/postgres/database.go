package postgres

import (
	"context"
	"errors"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.globerce.com/freedom-business/libs/shared-libs/databases/postgresql"
	"server/src/config"
	"server/src/repository"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Database struct {
	db     *pgxpool.Pool
	logger log.Logger

	remotePaymentRepository repository.RemotePaymentRepository
}

func New(ctx context.Context, cfg config.PostgresConfig, logger log.Logger) (repository.Storage, error) {
	db, err := postgresql.InitConnect(
		ctx,
		cfg.MaxConns,
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DatabaseName,
	)
	if err != nil {
		return nil, err
	}

	store := &Database{
		db:     db,
		logger: log.With(logger, "module", "postgres"),
	}

	if err = store.migrateUp(cfg); err != nil {
		return nil, err
	}

	return store, nil
}

func (d *Database) migrateUp(cfg config.PostgresConfig) error {
	dbURL := "postgres://" +
		cfg.Username + ":" +
		cfg.Password + "@" +
		cfg.Host + ":" +
		cfg.Port + "/" +
		cfg.DatabaseName +
		"?sslmode=disable"

	m, err := migrate.New(
		"file://schema/migration",
		dbURL,
	)
	if err != nil {
		_ = level.Error(d.logger).Log("migration error", err)
		return err
	}

	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			_ = level.Info(d.logger).Log("migration status", "no change")
			return nil // No error, just no change needed
		}
		_ = level.Error(d.logger).Log("migration error", err)
		return err
	}

	return nil
}

func (d *Database) RemotePayment() repository.RemotePaymentRepository {
	if d.remotePaymentRepository != nil {
		return d.remotePaymentRepository
	}

	d.remotePaymentRepository = &RemotePayment{
		db:     d.db,
		logger: log.With(d.logger, "repository", "remotePayment"),
	}

	return d.remotePaymentRepository
}
