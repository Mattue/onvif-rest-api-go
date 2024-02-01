package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mattue/onif-rest-api-go/processors"
)

func main() {
	router := gin.Default()

	version := "v1"

	//should be POST
	router.GET("/"+version+"/:camera/absolute-move", ptz_request_processors.AbsoluteMove)
	router.GET("/"+version+"/:camera/get-configurations", ptz_request_processors.GetConfigurations)
	router.GET("/"+version+"/:camera/get-presets", ptz_request_processors.GetPresets)
	router.GET("/"+version+"/:camera/get-preset-tours", ptz_request_processors.GetPresetTours)
	//should be POST
	router.GET("/"+version+"/:camera/goto-home-position", ptz_request_processors.GotoHomePosition)
	//should be POST
	router.GET("/"+version+"/:camera/goto-preset", ptz_request_processors.GotoPreset)

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}
