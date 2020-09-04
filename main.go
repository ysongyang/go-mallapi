package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golangPro/go-mallapi/app/models"
	"golangPro/go-mallapi/pkg/setting"
	"golangPro/go-mallapi/routers"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
	models.Setup()
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
