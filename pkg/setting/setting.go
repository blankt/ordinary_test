package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	PageSize int
}

var DatabaseSetting = &Database{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

func Setup() {
	config := viper.New()
	config.SetConfigName("app")
	config.SetConfigType("yaml")
	config.AddConfigPath("./conf")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("setting.Setup,fail to parse `conf/app.yaml`:%v", err)
	}
	err = config.UnmarshalKey("database", &DatabaseSetting)
	fmt.Println(DatabaseSetting)
	if err != nil {
		log.Fatalf("setting.Setup,fail to parse `conf/app.yaml`:%v", err)
	}
	err = config.UnmarshalKey("server", &ServerSetting)
	fmt.Println(ServerSetting)
	if err != nil {
		log.Fatalf("setting.Setup,fail to parse `conf/app.yaml`:%v", err)
	}
}
