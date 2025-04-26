package db

import (
	"context"
	"database/sql"
	"embed"
	"log"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
)

//go:embed migrations/*.sql
var migrationsFiles embed.FS

func ConnectDB() (*bun.DB, error) {
	dsn := os.Getenv("DB_DSN")

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	if err := RunMigrations(db); err != nil {
		log.Printf("failed to run migrations: %v", err)
		return nil, err
	}

	err := db.Ping()
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
	}
	return db, nil
}

// RunMigrations runs all .sql files
func RunMigrations(db *bun.DB) error {
	migrations := migrate.NewMigrations()
	if err := migrations.Discover(migrationsFiles); err != nil {
		return err
	}

	migrator := migrate.NewMigrator(db, migrations)
	if err := migrator.Init(context.Background()); err != nil {
		return err
	}

	_, err := migrator.Migrate(context.Background())
	if err != nil {
		return err
	}

	log.Println(" Migrations applied successfully")
	return nil
}

func InitDB() (*Conn, error) {
	conn := Conn{}

	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	conn.DB = db
	return &conn, nil
}
