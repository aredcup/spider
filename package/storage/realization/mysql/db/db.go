package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"redcup/spider/package/storage/options"
	"time"
)

var db *mysqlDB

type mysqlDB struct {
	engine *xorm.EngineGroup
	tables []interface{}
}

func InitEngine(options *options.StorageOptions) {

	var err error

	db.engine, err = xorm.NewEngineGroup("mysql", options.DataSourceName, xorm.RoundRobinPolicy())
	if err != nil {
		panic(err)
	}

	//设置打印日志
	db.engine.ShowSQL(true)

	//engine.SetMaxIdleConns()
	//engine.SetMaxIdleConns()

	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	//engine.SetDefaultCacher(cacher)

	if options.IsDropTables {
		err = db.engine.DropTables(db.tables...)
		if err != nil {
			return
		}
	}

	if options.IsSyncTables {
		err = db.engine.Sync2(db.tables...)
		if err != nil {
			return
		}
	}

	//engine.SetMaxIdleConns(config.GetMaxIdleConns())
	//engine.SetMaxOpenConns(config.GetMaxOpenConns())

	go func() {
		ticker := time.NewTicker(time.Second * time.Duration(options.ConnInterval))
		for range ticker.C {
			if err = db.engine.Ping(); err != nil {
				log.Printf("[Database] ping error :%s", err)
			}
		}
	}()
}

func AddTables(table ...interface{}) {
	db.tables = append(db.tables, table...)
}

func init() {
	db = &mysqlDB{
		tables: make([]interface{}, 0),
	}
}

func GetWriteDB() *xorm.Engine {
	return db.engine.Master()
}
