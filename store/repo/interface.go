package repo

import (
	"github.com/4lexvav/hit-the-goal/store/repo/postgres/lists"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres/projects"
)

type Repo interface {
	Projects() projects.DAO
	Lists() lists.DAO
}
