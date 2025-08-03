package database

import (
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MysqlConnect *gorm.DB
	mysqlOnce    sync.Once
)

func InitMysql(dsn string) {
	mysqlOnce.Do(func() {

		db, err := gorm.Open(mysql.Open(dsn))
		if err != nil {
			fmt.Printf("Lỗi kết nối tới database Mysql: %v", err)
		}

		MysqlConnect = db

		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}

		// ⚙️ Cấu hình connection pool
		sqlDB.SetMaxOpenConns(100)                 // số lượng connection tối đa tới MySQL
		sqlDB.SetMaxIdleConns(10)                  // số lượng connection giữ lại ở trạng thái idle
		sqlDB.SetConnMaxLifetime(30 * time.Minute) // thời gian sống tối đa của một connection
		sqlDB.SetConnMaxIdleTime(5 * time.Minute)  // thời gian tối đa connection có thể idle
	})
}
