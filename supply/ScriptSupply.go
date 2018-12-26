package supply

import (
	"log"

	"github.com/labstack/echo"
)

type ScriptSupplier struct {
	Logger *log.Logger `inject:""`
	Echo   *echo.Echo  `inject:""`
}

func (scriptSupplier *ScriptSupplier) Before() {
	scriptGroup := scriptSupplier.Echo.Group("/script")
	scriptGroup.GET("/test", scriptSupplier.Test)
}

func (scriptSupplier *ScriptSupplier) Test(c echo.Context) error {
	return nil
}