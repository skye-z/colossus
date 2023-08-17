package backend

import (
	"log"

	"github.com/skye-z/colossus/backend/model"
	"xorm.io/xorm"
)

func GetDBEngine() {
	engine := loadDBEngine()
	go initDatabase(engine)
}

func loadDBEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("sqlite", "./aircraft.store")
	if err != nil {
		panic(err)
	}
	return engine
}

func initDatabase(engine *xorm.Engine) {
	log.Println("[DB] start loading")
	err := engine.Sync2(new(model.Host))
	if err != nil {
		panic(err)
	}
	log.Println("[DB] loading completed")
}
