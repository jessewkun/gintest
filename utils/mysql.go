package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Mysql struct {
	// mysql dsn
	Dsn string `yaml:"dsn"`
	// 最大连接数
	MaxIdleConn int `yaml:"max_idle_conn"`
	// 最大空闲连接数
	MaxOpenConn int `yaml:"max_open_conn"`
}

var DB *gorm.DB

func InitDB(m *Mysql) {
	db, err := gorm.Open(mysql.Open(m.Dsn), &gorm.Config{
		// 打印所有执行的SQL
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名不加s
		},
	})
	if err != nil {
		fmt.Println("mysql connect error: " + err.Error())
		os.Exit(1)
	}

	sqlDB, err := db.DB()
	if err != nil {
		_ = sqlDB.Close()
		fmt.Println("mysql db error: " + err.Error())
		os.Exit(1)
	}

	sqlDB.SetMaxIdleConns(m.MaxIdleConn)
	sqlDB.SetMaxOpenConns(m.MaxOpenConn)

	// 获取sql配置情况
	sqlStats, _ := json.Marshal(sqlDB.Stats())
	fmt.Println(string(sqlStats))
	DB = db
}

func CloseDb() {
	sqlDB, err := DB.DB()
	if err != nil {
		_ = sqlDB.Close()
		fmt.Println("mysql close error: " + err.Error())
		os.Exit(1)
	}
	sqlDB.Close()
}
