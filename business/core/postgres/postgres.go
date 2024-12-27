package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	config "github.com/danipurwadi/db-backup-cli/foundation"
	"github.com/jackc/pgx/v5"
)

type Service interface {
	Backup(cfg config.Config) error
}

type PostgresCore struct {
	cfg *config.Config
}

func NewPostgresCore(cfg *config.Config) *PostgresCore {
	return &PostgresCore{
		cfg: cfg,
	}
}

func (p *PostgresCore) Connect() (*pgx.Conn, error) {
	slog.Info("initializing postgres...")
	conn, err := pgx.Connect(context.Background(), p.cfg.DbUrl)
	if err != nil {
		slog.Error(fmt.Sprintf("unable to connect to database %s \n", p.cfg.DbUrl), "err", err)
		return nil, err
	}
	return conn, nil
}

func (p *PostgresCore) Backup(conn *pgx.Conn) error {
	tableNames := p.cfg.Tables
	tables := strings.Split(tableNames, ",")
	export_dir := "output"

	for _, t := range tables {
		fileName := fmt.Sprintf("%s/%s_backup_%d.csv", export_dir, t, time.Now().UnixMilli())
		ef, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			slog.Error(fmt.Sprintf("error writing backup for table %s", t), "err", err)
			return err
		}
		defer ef.Close()

		err = exporter(conn, ef, t)
		if err != nil {
			slog.Info("failed to export", "err", err)
			break
		}
	}

	return nil
}

func exporter(conn *pgx.Conn, f *os.File, table string) error {
	res, err := conn.PgConn().CopyTo(context.Background(), f, fmt.Sprintf("COPY %s TO STDOUT DELIMITER ',' CSV HEADER", table))
	if err != nil {
		return fmt.Errorf("error exporting file: %+v", err)
	}

	slog.Info("successfully exported db", "rows affected", res.RowsAffected())
	return nil
}
