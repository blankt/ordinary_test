package main

import (
	"fmt"
	"net/http"
	"ordinary_test/model"
	"ordinary_test/pkg/setting"
	"ordinary_test/routers"
)

func init() {
	setting.Setup()
	model.Setup()
}

func main() {
	router := routers.InitRouter()

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	s := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	s.ListenAndServe()
}
