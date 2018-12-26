package supply

import (
	"github.com/facebookgo/inject"
)

type Injection struct {
	controllers []BaseSupplier
}

func (c *Injection) ComponentsName() string {
	return "Suppliers"
}

func (c *Injection) InjectingObjects() []*inject.Object {
	var userSupplier UserSupplier
	var scriptSupplier ScriptSupplier
	c.controllers = []BaseSupplier{
		&userSupplier,
		&scriptSupplier,
	}
	return []*inject.Object{
		{Value: &userSupplier},
		{Value: &scriptSupplier},
	}
}

func (c *Injection) After() error {
	for _, controller := range c.controllers {
		controller.Before()
	}
	return nil
}
