package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/Astemirdum/transactions/tx-balance/config"

	"github.com/Astemirdum/transactions/tx-balance/migrations"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

func NewDB(cfg *config.DB) (*sqlx.DB, error) {
	dsn := newDSN(cfg)
	if err := MigrateSchema(dsn); err != nil {
		return nil, err
	}

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func newDSN(cfg *config.DB) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.NameDB, cfg.Password)
}

func MigrateSchema(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	src, err := httpfs.New(http.FS(migrations.MigrationFiles), ".")
	if err != nil {
		return err
	}
	targetInstance, err := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: migrations.Label + "_migrations",
	})
	if err != nil {
		return fmt.Errorf("cannot create target db instance: %w", err)
	}

	m, err := migrate.NewWithInstance("<embed>", src, "postgres", targetInstance)
	if err != nil {
		return fmt.Errorf("cannot create migration instance: %w", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migrations failed: %w", err)
	}

	return targetInstance.Close()
}
