package main

import (
	"database/sql"
	"embed"

	"readcommend/internal/bootstrap"

	_ "github.com/jackc/pgx/v5"
	"github.com/pressly/goose/v3"
)

//go:embed scripts/*.sql
var embedMigrations embed.FS

func main() {
	// load yaml configuration
	config, err := bootstrap.LoadConfig()
	if err != nil {
		panic(err)
	}

	// setup database
	database, err := sql.Open("pgx", config.Database.URL)
	if err != nil {
		panic(err)
	}
	goose.SetBaseFS(embedMigrations)

	// setup goose
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	// run scripts
	if err := goose.Up(database, "scripts"); err != nil {
		panic(err)
	}
}
