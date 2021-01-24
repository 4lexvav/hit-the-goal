package comments

import (
	"sync"

	"github.com/4lexvav/hit-the-goal/store/repo"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres/comments"
)

type service struct {
	commentsDao comments.DAO
}

var (
	srv  service
	once = &sync.Once{}
)

func New() (_ Service, err error) {
	once.Do(func() {
		srv = service{commentsDao: repo.Get().Comments()}
	})

	return srv, err
}

func (srv service) WithTx(tx *postgres.DBQuery) Service {
	return service{commentsDao: repo.Get().Comments().WithTx(tx)}
}
