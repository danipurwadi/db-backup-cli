package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	config "github.com/danipurwadi/db-backup-cli/foundation"
	"github.com/jackc/pgx/v5"
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

func (p *PostgresClient) Initialise() (*pgx.Conn, error) {
	log.Println("initializing postgres...")
	conn, err := pgx.Connect(context.Background(), p.cfg.DbUrl)
	if err != nil {
		log.Println("unable to connec to database", p.cfg.DbUrl, err)
		return nil, err
	}
	return conn, nil
}

func (p *PostgresClient) Backup(conn *pgx.Conn) error {
	log.Println("performing backup for postgres table...")
	tables := []string{"table1", "table2", "table3"}

	import_dir := "/dir_to_import_from"
	for _, t := range tables {
		f, err := os.OpenFile(fmt.Sprintf("%s/table_%s.csv", import_dir, t), os.O_RDONLY, 0777)
		if err != nil {
			return err
		}
		f.Close()

		res, err := conn.PgConn().CopyFrom(context.Background(), f, fmt.Sprintf("COPY %s FROM STDIN DELIMITER '|' CSV HEADER", t))
		if err != nil {
			return err
		}
		fmt.Println("==> import rows affected:", res.RowsAffected())
	}

	return nil
}
