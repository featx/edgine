package depend

import "github.com/facebookgo/inject"

type Injection struct {
}

func (c *Injection) ComponentsName() string {
	return "DataAccess"
}

func (c *Injection) InjectingObjects() []*inject.Object {
	return []*inject.Object{
		{Value: &UserDataAccessor{}},
		{Value: &ScriptDataAccessor{}},
	}
}

func (c *Injection) After() error {
	return nil
}
