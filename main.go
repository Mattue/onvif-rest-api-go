package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mattue/onif-rest-api-go/processors/ptz"
)

func main() {
	router := gin.Default()

	version := "v1"

	//should be POST
	router.GET("/"+version+"/:camera/absolute-move", ptz.AbsoluteMove)
	router.GET("/"+version+"/:camera/get-configurations", ptz.GetConfigurations)
	router.GET("/"+version+"/:camera/get-presets", ptz.GetPresets)
	router.GET("/"+version+"/:camera/get-preset-tours", ptz.GetPresetTours)
	//should be POST
	router.GET("/"+version+"/:camera/goto-home-position", ptz.GotoHomePosition)
	//should be POST
	router.GET("/"+version+"/:camera/goto-preset", ptz.GotoPreset)

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}
