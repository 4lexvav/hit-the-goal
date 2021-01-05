package postgres

import (
	"sync"
	"time"

	"github.com/4lexvav/hit-the-goal/config"
	"github.com/go-pg/pg/v9"
)

type Postgres struct {
	conn *pg.DB
}

var (
	postgres *Postgres
	once     = &sync.Once{}
)

func Load(cfg config.Postgres, lgr DBLogger) error {
	once.Do(func() {
		db := pg.Connect(&pg.Options{
			Addr:         cfg.Host + ":" + cfg.Port,
			User:         cfg.User,
			Password:     cfg.Password,
			Database:     cfg.Database,
			PoolSize:     cfg.PoolSize,
			MaxRetries:   cfg.MaxRetries,
			ReadTimeout:  time.Duration(cfg.ReadTimeout),
			WriteTimeout: time.Duration(cfg.WriteTimeout),
		})

		db.AddQueryHook(dbLogger{logger: lgr})
		postgres = &Postgres{conn: db}
	})

	return postgres.Ping()
}

func GetDB() *Postgres {
	return postgres
}

func (p Postgres) Ping() error {
	var n int

	_, err := p.conn.QueryOne(pg.Scan(&n), "SELECT 1")

	return err
}
