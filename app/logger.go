package app

import (
	"log"
	"os"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	glog "github.com/labstack/gommon/log"
)

type LoggerBus struct {
	base *log.Logger
}

func base() *log.Logger {
	return log.New(os.Stdout, "[BUS]", log.Ldate|log.Lmicroseconds)
}

func (loggerBus *LoggerBus) init(logFileName string) *os.File {
	if loggerBus.base == nil {
		loggerBus.base = base()
	}
	logFile, err := os.Create(logFileName)
	if err != nil {
		loggerBus.base.Fatalf("Create log file [%s] for PIED error, cause: %s", logFileName, err.Error())
		return nil
	}
	return logFile
}

func (loggerBus *LoggerBus) EchoUse() *glog.Logger {
	logFile := loggerBus.init("log/echo.log")
	var glogger = glog.New("[ECH]")
	//glogger.SetHeader("")
	glogger.SetOutput(logFile)
	return glogger
}

func (loggerBus *LoggerBus) BusinessUse() *log.Logger {
	logFile := loggerBus.init("log/pied.log")
	return log.New(logFile, "[PID]", log.Ldate|log.Lmicroseconds)
}

func (loggerBus *LoggerBus) PersistenceUse() core.ILogger {
	logFile := loggerBus.init("log/sql.log")
	return xorm.NewSimpleLogger2(logFile, "[SQL]", log.Ldate|log.Lmicroseconds)
}
