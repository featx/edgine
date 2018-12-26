package supply

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type UserSupplier struct {
	Logger *log.Logger `inject:""`
	Echo   *echo.Echo  `inject:""`
}

func (userSupplier *UserSupplier) Before() {
	userGroup := userSupplier.Echo.Group("/user")
	userGroup.GET("/test", userSupplier.Test)
}

func (userSupplier *UserSupplier) Test(c echo.Context) error {
	userSupplier.Logger.Printf("Start time %s", time.Now())
	req := c.Request()
	format := `
    <code>
      Protocol: %s<br>
      Host: %s<br>
      Remote Address: %s<br>
      Method: %s<br>
      Path: %s<br>
    </code>
  `
	userSupplier.Logger.Printf("End time %s", time.Now())
	return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Referer(), req.Host, req.RemoteAddr, req.Method, req.URL.Path))
}
