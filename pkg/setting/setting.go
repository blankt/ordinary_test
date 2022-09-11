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
	RunMode      string        `mapstructure:"run_mode"`
	HttpPort     int           `mapstructure:"http_port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

var ServerSetting = &Server{}

func Setup() {
	viper.SetConfigName("app")
	viper.SetConfigType("ini")
	viper.AddConfigPath("D:\\goProject\\ordinary_test\\conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("setting.Setup,fail to parse `conf/app.ini`:%v", err)
	}
	err = viper.Unmarshal(&DatabaseSetting)
	fmt.Println(DatabaseSetting)
	if err != nil {
		log.Fatalf("setting.Setup,fail to parse `conf/app.ini`:%v", err)
	}
	err = viper.Unmarshal(&ServerSetting)
	fmt.Println(ServerSetting)
	if err != nil {
		log.Fatalf("setting.Setup,fail to parse `conf/app.ini`:%v", err)
	}
}
