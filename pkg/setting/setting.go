package setting

import (
	"github.com/spf13/viper"
	"log"
)

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

var DatabaseSetting = &Database{}

func Setup() {
	viper.SetConfigName("app")
	viper.SetConfigType("ini")
	viper.AddConfigPath("/conf/app.ini")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("setting.Setup,fail to parse `conf/app.ini`:%v", err)
	}
	err = viper.Unmarshal(&DatabaseSetting)
	if err != nil {
		log.Fatalf("setting.Setup,fail to parse `conf/app.ini`:%v", err)
	}
}
