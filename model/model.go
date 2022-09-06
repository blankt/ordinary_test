package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"ordinary_test/pkg/setting"
)

var db *gorm.DB

func Setup() {
	var err error
	if db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.DBName,
		setting.DatabaseSetting.Port)), &gorm.Config{}); err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}
