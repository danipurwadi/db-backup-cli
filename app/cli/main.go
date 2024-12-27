package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"slices"

	"github.com/danipurwadi/db-backup-cli/business/core/postgres"
	config "github.com/danipurwadi/db-backup-cli/foundation"
)

type dbType string

// db types
const (
	pg    dbType = "postgres"
	mysql dbType = "mysql"
	mongo dbType = "mongo"
)

func AllDbs() []dbType {
	return []dbType{pg, mysql, mongo}
}

func main() {
	slog.Info("starting CLI application...")

	config := config.Config{}
	flag.StringVar(&config.DbName, "d", "", "db name")
	flag.StringVar(&config.DbType, "t", "", "database type")
	flag.StringVar(&config.DbUrl, "h", "", "host url")
	flag.StringVar(&config.Username, "u", "", "username")
	flag.StringVar(&config.Password, "w", "", "password")
	flag.StringVar(&config.Tables, "n", "", "table names")
	flag.Parse()

	err := validateConfig(&config)
	if err != nil {
		slog.Error("failed to startup application due to invalid config", "err", err)
		return
	}
	switch dt := dbType(config.DbType); dt {
	case pg:
		service := postgres.NewPostgresCore(&config)
		conn, err := service.Connect()
		if err != nil {
			slog.Error("failed to connect to postgres db", "err", err)
			return
		}

		slog.Info("successfully connected to postgres!")
		err = service.Backup(conn)
		if err != nil {
			slog.Error("failed to perform backup for postgres", "err", err)
			return
		}

		slog.Info("successfully perform backup!")
	}
}

func validateConfig(cfg *config.Config) error {
	if cfg.DbName == "" {
		log.Printf("missing db name in flag parameter. Usage -d <db_name>")
		return fmt.Errorf("invalid db name")
	}

	if !slices.Contains(AllDbs(), dbType(cfg.DbType)) {
		log.Printf("invalid db type in flag parameter. Valid db types %+v\n", AllDbs())
		return fmt.Errorf("invalid db type")
	}

	if cfg.DbUrl == "" {
		log.Printf("missing db host url in flag parameter. Usage -h <host_url>")
		return fmt.Errorf("invalid db host url")
	}

	if cfg.Username == "" {
		log.Printf("missing username in flag parameter. Usage -u <username>")
		return fmt.Errorf("invalid username")
	}

	if cfg.Password == "" {
		log.Printf("missing password in flag parameter. Usage -w <password>")
		return fmt.Errorf("invalid password")
	}
	return nil
}
