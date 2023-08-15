package main

import (
	"flag"
	"fmt"

	"github.com/bagusyanuar/go-yousee/cmd/database/scheme"
	"github.com/bagusyanuar/go-yousee/config"
)

func main() {
	cfg, err := config.NewConfig(".env")
	if err != nil {
		panic(err)
	}

	db, err := config.NewMySQLConnection(&cfg.MySQL)

	if err != nil {
		panic(err)
	}

	cmd := flag.String("m", "", "unsupport command")
	flag.Parse()
	command := *cmd

	switch command {
	case "fresh":
		scheme.Drop(db)
		scheme.Migrate(db)
		fmt.Println("successfully fresh database")
	case "seed":
		// seeder.Seed(db)
		fmt.Println("successfully seed database")
	default:
		scheme.Migrate(db)
		fmt.Println("successfully migrating database")
		return
	}
	scheme.Migrate(db)
	fmt.Println("successfully migrating database")
}
