package repo

import (
	"sync"

	"github.com/4lexvav/hit-the-goal/store/repo/postgres/lists"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres/projects"
)

var (
	repo postgresRepo
	once = &sync.Once{}
)

type postgresRepo struct {
	projectsDao projects.DAO
	listsDao    lists.DAO
}

func Get() Repo {
	return repo
}

func Load() (err error) {
	once.Do(func() {
		repo = postgresRepo{
			projectsDao: projects.NewProjectsDao(),
			listsDao:    lists.NewListDao(),
		}
	})

	return err
}

func (r postgresRepo) Projects() projects.DAO {
	return r.projectsDao
}

func (r postgresRepo) Lists() lists.DAO {
	return r.listsDao
}
