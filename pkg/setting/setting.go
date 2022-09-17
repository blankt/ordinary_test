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
}

var DatabaseSetting = &Database{}

type Server struct {
	RunMode      string        `mapstructure:"runmode"`
	HttpPort     int           `mapstructure:"httpport"`
	ReadTimeout  time.Duration `mapstructure:"readtimeout"`
	WriteTimeout time.Duration `mapstructure:"writetimeout"`
}

var ServerSetting = &Server{}

func Setup() {
	config := viper.New()
	config.SetConfigName("app")
	config.SetConfigType("yaml")
	config.AddConfigPath("conf")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("setting.Setup,fail to parse `conf/app.yaml`:%v", err)
	}
	err = config.Unmarshal(&DatabaseSetting)
	fmt.Println(DatabaseSetting)
	if err != nil {
		log.Fatalf("setting.Setup,fail to parse `conf/app.yaml`:%v", err)
	}
	err = config.Unmarshal(&ServerSetting)
	fmt.Println(ServerSetting)
	if err != nil {
		log.Fatalf("setting.Setup,fail to parse `conf/app.yaml`:%v", err)
	}
}
