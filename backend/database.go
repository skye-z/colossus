package backend

import (
	"fmt"
	"log"
	"os"

	"github.com/skye-z/colossus/backend/model"
	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

func GetDBEngine() *xorm.Engine {
	engine := loadDBEngine()
	go initDatabase(engine)
	return engine
}

func loadDBEngine() *xorm.Engine {
	// 获取用户配置文件目录
	configDir, _ := os.UserConfigDir()
	configDir = fmt.Sprintf("%s/Colossus", configDir)
	// 判断应用目录是否存在
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		// 目录不存在,创建目录
		os.Mkdir(configDir, os.ModePerm)
	}
	// 组装数据库存储路径
	dbDir := fmt.Sprintf("%s/%s", configDir, "colossus.store")
	log.Println(dbDir)
	log.Println("Database path: " + dbDir)
	engine, err := xorm.NewEngine("sqlite", dbDir)
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
