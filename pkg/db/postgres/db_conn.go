package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/qsmsoft/blog/config"
	"log"
	"time"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

type Database struct {
	Conn *sqlx.DB
}

func NewDatabase(cfg *config.Config) (*Database, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.Postgres.PostgresqlUser, cfg.Postgres.PostgresqlPassword, cfg.Postgres.PostgresqlHost, cfg.Postgres.PostgresqlPort, cfg.Postgres.PostgresqlDBName, cfg.Postgres.PostgresqlSSLMode)

	db, err := sqlx.Connect(cfg.Postgres.PostgresqlDriver, dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Database{Conn: db}, nil
}

func (db *Database) Close() {
	if err := db.Conn.Close(); err != nil {
		log.Printf("failed to close database connection: %v", err)
	} else {
		log.Println("database connection closed")
	}
}
