package postgres

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/4lexvav/hit-the-goal/config"
	"github.com/4lexvav/hit-the-goal/logger"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once = &sync.Once{}
)

func Load(cfg config.Postgres) error {
	once.Do(func() {
		dataSourceName := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", cfg.Database, cfg.User, cfg.Password, cfg.Host, cfg.Port)
		newDB, err := sql.Open("postgres", dataSourceName)

		if err != nil {
			logger.Get().Fatalw("Cannot open DB connection", "error", err)
		}

		newDB.SetMaxOpenConns(cfg.PoolSize)

		db = newDB
	})

	return db.Ping()
}

func GetDB() *sql.DB {
	return db
}
