package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golangPro/golang-mallapi/app/models"
	"golangPro/golang-mallapi/pkg/gredis"
	"golangPro/golang-mallapi/pkg/logging"
	"golangPro/golang-mallapi/pkg/setting"
	"golangPro/golang-mallapi/pkg/util"
	"golangPro/golang-mallapi/routers"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

func main() {

	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
