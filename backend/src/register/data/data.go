package data

import (
	"time"

	"EDA.Project.ERP.backend.register/config"
	"EDA.Project.ERP.backend.register/tools/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open(mysql.Open(config.MySQLdb), &gorm.Config{})
	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect error": err.Error()})
	}
	if Db.Error != nil {
		logger.Error(map[string]interface{}{"database error": Db.Error})
	}
	sqlDB, _ := Db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}
