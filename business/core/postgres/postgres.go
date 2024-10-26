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
		log.Printf("unable to connect to database %s %s \n", p.cfg.DbUrl, err)
		return nil, err
	}
	return conn, nil
}

func (p *PostgresClient) Backup(conn *pgx.Conn) error {
	log.Println("performing backup for postgres table...")
	tables := []string{"example_table"}

	export_dir := "output"
	for _, t := range tables {
		ef, err := os.OpenFile(fmt.Sprintf("%s/table_%s.csv", export_dir, t), os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			fmt.Println("error opening file:", err)
			return err
		}
		defer ef.Close()

		err = exporter(conn, ef, t)
		if err != nil {
			log.Println("failed to export", err)
			break
		}
	}

	return nil
}

func exporter(conn *pgx.Conn, f *os.File, table string) error {
	res, err := conn.PgConn().CopyTo(context.Background(), f, fmt.Sprintf("COPY %s TO STDOUT DELIMITER '|' CSV HEADER", table))
	if err != nil {
		return fmt.Errorf("error exporting file: %+v", err)
	}
	fmt.Println("==> export rows affected:", res.RowsAffected())
	return nil
}
