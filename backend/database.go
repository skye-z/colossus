package backend

import (
	"fmt"
	"log"
	"os"

	"github.com/skye-z/colossus/backend/model"
	"xorm.io/xorm"
)

func GetDBEngine() *xorm.Engine {
	engine := loadDBEngine()
	go initDatabase(engine)
	return engine
}

func loadDBEngine() *xorm.Engine {
	configDir, _ := os.UserConfigDir()
	engine, err := xorm.NewEngine("sqlite", fmt.Sprintf("%s/%s", configDir, "colossus.store"))
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
