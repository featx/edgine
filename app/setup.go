package app

import (
	"log"

	"github.com/facebookgo/inject"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/featx/edgine/depend"
	"github.com/featx/edgine/hinge"
	"github.com/featx/edgine/supply"
)

type Edgine struct {
	logger     *log.Logger
	httpEngine *echo.Echo
}

type Injectable interface {
	ComponentsName() string

	InjectingObjects() []*inject.Object

	After() error
}

func (app *Edgine) Init() {
	var loggerBus LoggerBus
	var graph inject.Graph

	//Initial Edgine inside dependencies
	app.logger = loggerBus.BusinessUse()

	app.httpEngine = echo.New()

	app.httpEngine.Logger = loggerBus.EchoUse()
	//app.httpEngine.DefaultHTTPErrorHandler()
	app.httpEngine.Use(middleware.LoggerWithConfig(loggerBus.HttpConfig()))
	app.httpEngine.Use(middleware.Recover())

	err := graph.Provide(
		&inject.Object{Value: app.logger},
		&inject.Object{Value: app.httpEngine},
		&inject.Object{Value: InitPersistence(loggerBus.PersistenceUse())},
	)

	if err != nil {
		loggerBus.base.Fatalf("Components intial unsuccessfully, cause: %s", err.Error())
	}
	injections := []Injectable{
		&supply.Injection{},
		&hinge.Injection{},
		&depend.Injection{},
	}
	for _, injection := range injections {
		app.registerGraph(&graph, injection)
	}
	if err := graph.Populate(); err != nil {
		loggerBus.base.Fatalf("Components assembled unsuccessfully, cause: %s", err.Error())
	}
	for _, injection := range injections {
		if err := injection.After(); err != nil {
			loggerBus.base.Fatalf("Injection comp[%s] unsuccessfully, cause: %s", injection.ComponentsName(),
				err.Error())
		}
	}
}

func (app *Edgine) Launch() {
	err := app.httpEngine.Start(":8080")
	app.logger.Fatalf("Launch app error, cause: %s", err.Error())
}

func (app *Edgine) registerGraph(graph *inject.Graph, s Injectable) {
	err := graph.Provide(s.InjectingObjects()...)
	if err != nil {
		app.logger.Fatalf("Components [%s] Intial unsuccessfully, cause: %s", s.ComponentsName(), err.Error())
	}
}
