package depend

import "github.com/go-xorm/xorm"

type ScriptDataAccess interface {
}

type ScriptDataAccessor struct {
	Engine *xorm.Engine `inject:""`
}
