package projects

import (
	"sync"

	"github.com/4lexvav/hit-the-goal/store/repo"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres/projects"
)

type service struct {
	projectsDao projects.DAO
}

var (
	srv  service
	once = &sync.Once{}
)

func New() (_ Service, err error) {
	once.Do(func() {
		srv = service{projectsDao: repo.Get().Projects()}
	})

	return srv, err
}
