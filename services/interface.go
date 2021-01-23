package services

import (
	"github.com/4lexvav/hit-the-goal/services/lists"
	"github.com/4lexvav/hit-the-goal/services/projects"
	"github.com/4lexvav/hit-the-goal/services/tasks"
)

type Service interface {
	Projects() projects.Service
	Lists() lists.Service
	Tasks() tasks.Service
}
