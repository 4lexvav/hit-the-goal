package postgres

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/4lexvav/hit-the-goal/config"
	"github.com/4lexvav/hit-the-goal/logger"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

var (
	postgres *Postgres
	once     = &sync.Once{}
)

func Load(cfg config.Postgres) error {
	once.Do(func() {
		dataSourceName := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=%s", cfg.Database, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Sslmode)
		logger.Get().Infow("Connection", "conn", dataSourceName)
		newDB, err := sql.Open("postgres", dataSourceName)

		if err != nil {
			logger.Get().Fatalw("Cannot open DB connection", "error", err)
		}

		newDB.SetMaxOpenConns(cfg.PoolSize)

		postgres = &Postgres{
			db: newDB,
		}
	})

	return postgres.db.Ping()
}

func GetDB() *Postgres {
	return postgres
}
