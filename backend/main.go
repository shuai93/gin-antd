package main

import (
	"backend/base"
	"backend/conf"
	"backend/models"
	"backend/utils/gredis"
	"backend/utils/setting"
	"fmt"
	"net/http"
)

func main() {
	conf.Init()
	setting.Setup()
	models.Setup()
	_ = gredis.Setup()

	router := base.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()
}
