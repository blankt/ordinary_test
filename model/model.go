package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"ordinary_test/pkg/setting"
)

var db *gorm.DB

var DatabaseAutoMigrate = []interface{}{
	&Article{},
	&Auth{},
	&Tag{},
}

func Setup() {
	var err error
	if db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.DBName,
		setting.DatabaseSetting.Port)),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: logger.Default.LogMode(logger.Info),
		},
	); err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	db.AutoMigrate(DatabaseAutoMigrate...)
}
