package main

import (
	"flag"
	"fmt"
)

type config struct {
	dbname   string
	dbtype   string
	dburl    string
	username string
	password string
}

func main() {
	fmt.Println("starting application...")

	config := config{}
	flag.StringVar(&config.dbname, "n", "", "db name")
	flag.StringVar(&config.dbtype, "t", "", "database type")
	flag.StringVar(&config.dburl, "d", "", "db url")
	flag.StringVar(&config.username, "u", "", "user name")
	flag.StringVar(&config.password, "w", "", "password")
	flag.Parse()

	fmt.Printf("loaded configs %+v \n", config)
	
}
