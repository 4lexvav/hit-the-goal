package tasks

import (
	"sync"

	"github.com/4lexvav/hit-the-goal/store/repo"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres/tasks"
)

type service struct {
	tasksDao tasks.DAO
}

var (
	srv  service
	once = &sync.Once{}
)

func New() (_ Service, err error) {
	once.Do(func() {
		srv = service{tasksDao: repo.Get().Tasks()}
	})

	return srv, err
}

func (srv service) WithTx(tx *postgres.DBQuery) Service {
	return service{tasksDao: repo.Get().Tasks().WithTx(tx)}
}
