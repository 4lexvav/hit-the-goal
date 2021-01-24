package services

import (
	"sync"

	"github.com/4lexvav/hit-the-goal/services/comments"
	"github.com/4lexvav/hit-the-goal/services/lists"
	"github.com/4lexvav/hit-the-goal/services/projects"
	"github.com/4lexvav/hit-the-goal/services/tasks"
)

var (
	srvRepo serviceRepo
	once    = &sync.Once{}
)

type serviceRepo struct {
	projectsSrv projects.Service
	listsSrv    lists.Service
	tasksSrv    tasks.Service
	commentsSrv comments.Service
}

func (srv serviceRepo) Projects() projects.Service {
	return srvRepo.projectsSrv
}

func (srv serviceRepo) Lists() lists.Service {
	return srvRepo.listsSrv
}

func (srv serviceRepo) Tasks() tasks.Service {
	return srvRepo.tasksSrv
}

func (srv serviceRepo) Comments() comments.Service {
	return srvRepo.commentsSrv
}

func Get() Service {
	return srvRepo
}

func Load() (err error) {
	once.Do(func() {
		projectsSrv, e := projects.New()
		if e != nil {
			err = e
			return
		}

		listsSrv, e := lists.New()
		if e != nil {
			err = e
			return
		}

		tasksSrv, e := tasks.New()
		if e != nil {
			err = e
			return
		}

		commentsSrv, e := comments.New()
		if e != nil {
			err = e
			return
		}

		srvRepo = serviceRepo{
			projectsSrv: projectsSrv,
			listsSrv:    listsSrv,
			tasksSrv:    tasksSrv,
			commentsSrv: commentsSrv,
		}
	})

	return err
}
