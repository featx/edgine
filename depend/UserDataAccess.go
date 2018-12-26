package depend

import "github.com/go-xorm/xorm"

type UserDataAccess interface {
}

type UserDataAccessor struct {
	Engine *xorm.Engine `inject:""`
}
