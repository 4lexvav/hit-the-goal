package services

import (
	"github.com/4lexvav/hit-the-goal/services/lists"
	"github.com/4lexvav/hit-the-goal/services/projects"
)

type Service interface {
	Projects() projects.Service
	Lists() lists.Service
}
