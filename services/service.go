package services

import (
	"sync"

	"github.com/4lexvav/hit-the-goal/services/projects"
)

var (
	srvRepo serviceRepo
	once    = &sync.Once{}
)

type serviceRepo struct {
	projectsSrv projects.Service
}

func (srv serviceRepo) Projects() projects.Service {
	return srvRepo.projectsSrv
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

		srvRepo = serviceRepo{projectsSrv: projectsSrv}
	})

	return err
}
