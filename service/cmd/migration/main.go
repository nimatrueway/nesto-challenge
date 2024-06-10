package main

import (
	"database/sql"
	"embed"
	_ "github.com/jackc/pgx/v5"
	"github.com/pressly/goose/v3"
	"readcommend/internal/bootstrap"
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
	db, err := sql.Open("pgx", config.Database.Dns)
	if err != nil {
		panic(err)
	}

	// setup goose
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	// run scripts
	if err := goose.Up(db, "scripts"); err != nil {
		panic(err)
	}
}
