package repo

import (
	"sync"

	"github.com/4lexvav/hit-the-goal/store/repo/postgres/lists"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres/projects"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres/tasks"
)

var (
	repo postgresRepo
	once = &sync.Once{}
)

type postgresRepo struct {
	projectsDao projects.DAO
	listsDao    lists.DAO
	tasksDao    tasks.DAO
}

func Get() Repo {
	return repo
}

func Load() (err error) {
	once.Do(func() {
		repo = postgresRepo{
			projectsDao: projects.NewProjectsDao(),
			listsDao:    lists.NewListDao(),
			tasksDao:    tasks.NewTaskDao(),
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

func (r postgresRepo) Tasks() tasks.DAO {
	return r.tasksDao
}
