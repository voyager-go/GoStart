package mysql

import (
	"fmt"
	"go-start/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var Conn *gorm.DB

func NewMysql() {
	cfg := config.Cfg.DataBase
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	var err error
	Conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 是否设置单数表名，设置为 是
		},
	})
	if err != nil {
		panic(fmt.Errorf("无法连接数据库，请先检查数据库配置信息，错误详情为: %s", err.Error()))
	}
	// GORM 使用 database/sql 维护连接池
	sqlDb, _ := Conn.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDb.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDb.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDb.SetConnMaxLifetime(time.Hour)
}
