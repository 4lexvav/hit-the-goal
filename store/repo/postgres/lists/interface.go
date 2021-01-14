package lists

import "github.com/4lexvav/hit-the-goal/models"

type DAO interface {
	Insert(list models.List) (_ models.List, err error)
}
