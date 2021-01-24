package repo

import (
	"github.com/4lexvav/hit-the-goal/store/repo/postgres/comments"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres/lists"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres/projects"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres/tasks"
)

type Repo interface {
	Projects() projects.DAO
	Lists() lists.DAO
	Tasks() tasks.DAO
	Comments() comments.DAO
}
