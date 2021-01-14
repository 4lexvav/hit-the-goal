package services

import (
	"sync"

	"github.com/4lexvav/hit-the-goal/services/lists"
	"github.com/4lexvav/hit-the-goal/services/projects"
)

var (
	srvRepo serviceRepo
	once    = &sync.Once{}
)

type serviceRepo struct {
	projectsSrv projects.Service
	listsSrv    lists.Service
}

func (srv serviceRepo) Projects() projects.Service {
	return srvRepo.projectsSrv
}

func (srv serviceRepo) Lists() lists.Service {
	return srvRepo.listsSrv
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

		srvRepo = serviceRepo{
			projectsSrv: projectsSrv,
			listsSrv:    listsSrv,
		}
	})

	return err
}
