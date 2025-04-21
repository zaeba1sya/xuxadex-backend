package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/xuxadex/backend-mvp-main/config"
	"github.com/xuxadex/backend-mvp-main/pkg/logger"
)

const (
	migrationsDir = "db/migrations"
)

func main() {
	up := flag.Bool("up", false, "Apply all up migrations")
	down := flag.Bool("down", false, "Apply down migrations (rollback latest)")
	status := flag.Bool("status", false, "Show migration status")

	flag.Parse()

	cfg := config.InitConfig()

	log := logger.NewApiLogger(cfg)
	log.InitLogger()

	var sslMode string = "disable"

	if cfg.DB.SSL {
		sslMode = "require"
	}

	dbURL := fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.DB.Dialect, cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, sslMode,
	)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	defer db.Close()

	err = goose.SetDialect(cfg.DB.Dialect)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}

	switch {
	case *up:
		fmt.Println("Applying up migrations...")
		if err := goose.Up(db, migrationsDir); err != nil {
			log.Fatalf("Failed to apply up migrations: %v", err)
		}
	case *down:
		fmt.Println("Rolling back last migration...")
		if err := goose.DownTo(db, migrationsDir, -1); err != nil {
			log.Fatalf("Failed to apply down migration: %v", err)
		}
	case *status:
		fmt.Println("Checking migration status...")
		if err := goose.Status(db, migrationsDir); err != nil {
			log.Fatalf("Failed to check migration status: %v", err)
		}
	default:
		fmt.Println("Usage: go run ./cmd/migrations/migrations.go --up|--down|--status")
		os.Exit(1)
	}
}
