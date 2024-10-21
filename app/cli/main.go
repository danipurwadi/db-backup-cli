package main

import (
	"flag"
	"fmt"
	"log"
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
	log.Println("starting CLI application...")

	config := config.Config{}
	flag.StringVar(&config.DbName, "d", "", "db name")
	flag.StringVar(&config.DbType, "t", "", "database type")
	flag.StringVar(&config.DbUrl, "h", "", "host url")
	flag.StringVar(&config.Username, "u", "", "username")
	flag.StringVar(&config.Password, "w", "", "password")
	flag.Parse()

	err := validateConfig(&config)
	if err != nil {
		log.Fatal("failed to startup application due to invalid config", err)
	}
	log.Printf("successfully loaded configs %+v \n", config)

	switch dt := dbType(config.DbType); dt {
	case pg:
		service := postgres.NewPostgresClient(&config)
		err := service.Backup()
		if err != nil {
			log.Fatal("failed to perform backup for postgres", err)
		}
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
