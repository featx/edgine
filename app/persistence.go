package app

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

func InitPersistence(logger core.ILogger) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/pied?charset=utf8mb4")
	if err != nil {
		logger.Errorf("Initialize persistence engine Error", err)
		return nil
	}

	engine.SetLogger(logger)
	engine.ShowSQL(true)
	return engine
}
