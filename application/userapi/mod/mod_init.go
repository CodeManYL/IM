package users_model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

func InitModEngine(dbName,dbAddress string) (err error){
	engine, err = xorm.NewEngine(dbName, dbAddress)
	if err != nil {
		return
	}
	menber := &Users{}
	menber.Token = "2222"
	_, _ = engine.Insert(menber)
	return nil
}
