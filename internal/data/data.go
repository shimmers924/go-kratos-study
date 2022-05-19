package data

import (
	"kratos-realworld/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	// gorm.Model
	// Code  string
	// Price uint
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.01:3306)/realworld?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{PrepareStmt: true})
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})  //sqlite 方式
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema  生产环境不建议这么干单独搞一套数据迁移
	// db.AutoMigrate(&Data{})

	// Create
	// db.Create(&Data{Code: "D42", Price: 100})
	return db
}
