package main

import (
	"log"

	"encore.app/url"
	"encore.dev/storage/sqldb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// List of models to auto-migrate.
var models = []any{
	&url.URL{},
}

func main() {
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: gormdb.Stdlib()}), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.AutoMigrate(models...); err != nil {
		log.Fatalln(err)
	}
}

// Create a temporary database to run auto-migrate against.
var gormdb = sqldb.NewDatabase("gorm-url", sqldb.DatabaseConfig{
	Migrations: "./placeholder-migrations",
})
