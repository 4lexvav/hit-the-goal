package lists

import (
	"sync"

	"github.com/4lexvav/hit-the-goal/store/repo"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres/lists"
)

type service struct {
	listsDao lists.DAO
}

var (
	srv  service
	once = &sync.Once{}
)

func New() (_ Service, err error) {
	once.Do(func() {
		srv = service{listsDao: repo.Get().Lists()}
	})

	return srv, err
}
