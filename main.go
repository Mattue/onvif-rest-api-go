package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hooklift/gowsdl/soap"
	"log"
	"mattue/onif-rest-api-go/onvif_wsdl2go"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()

	router.GET("/v1/:camera/:operation", processRequest)

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func processRequest(context *gin.Context) {

	client := soap.NewClient("***REMOVED***",
		soap.WithTimeout(time.Second*5),
	)

	service := onvif_wsdl2go.NewPTZ(client)

	var err error
	var reply any

	path := context.Param("operation")
	switch path {
	case "get-configurations":
		reply, err = service.GetConfigurations(&onvif_wsdl2go.GetConfigurations{})
	case "goto-preset":
		presetToken := onvif_wsdl2go.ReferenceToken(context.Query("preset-token"))
		profileToken := onvif_wsdl2go.ReferenceToken(context.Query("profile-token"))

		reply, err = service.GotoPreset(&onvif_wsdl2go.GotoPreset{PresetToken: &presetToken, ProfileToken: &profileToken})
	default:
		log.Println("Nothing found")
		return
	}

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusInternalServerError, "{err:err}")
		return
	}

	//log.Println(reply.PTZConfiguration[0].Name)
	log.Println(reply)

	context.IndentedJSON(http.StatusOK, "{123:123}")
}
