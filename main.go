package main

import (
	"net/http"
	"ordinary_test/pkg/setting"
	"ordinary_test/routers"
	"time"
)

func init() {
	setting.Setup()
	//model.Setup()
}

func main() {
	router := routers.InitRouter()

	//readTimeout := setting.ServerSetting.ReadTimeout
	//writeTimeout := setting.ServerSetting.WriteTimeout
	//endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	//maxHeaderBytes := 1 << 20
	s := &http.Server{
		Addr:           ":8000",
		Handler:        router,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
