package postgres

import (
	"log"

	config "github.com/danipurwadi/db-backup-cli/foundation"
)

type Service interface {
	Backup(cfg config.Config) error
}

type PostgresClient struct {
	cfg *config.Config
}

func NewPostgresClient(cfg *config.Config) *PostgresClient {
	return &PostgresClient{
		cfg: cfg,
	}
}

func (p *PostgresClient) Backup() error {
	log.Println("performing backup for postgres...")

	return nil
}
