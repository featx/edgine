package hinge

import "github.com/facebookgo/inject"

type Injection struct {
}

func (c *Injection) ComponentsName() string {
	return "Hinge"
}

func (c *Injection) InjectingObjects() []*inject.Object {
	return []*inject.Object{
		{Value: &UserHinging{}},
		{Value: &ScriptHinging{}},
	}
}

func (c *Injection) After() error {
	return nil
}
