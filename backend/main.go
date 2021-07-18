package main

import (
	"backend/conf"
	"backend/models"
	"backend/routers"
	"backend/utils/gredis"
	"backend/utils/setting"
	"fmt"
	"net/http"
)

func main() {
	conf.Init()
	setting.Setup()
	models.Setup()
	gredis.Setup()

	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
