package ptz_request_processors

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"log"
	"mattue/onif-rest-api-go/errorsProcessors"
	"mattue/onif-rest-api-go/onvif_wsdl2go"
	"net/http"
)

func AbsoluteMove(context *gin.Context) {
	service := getService(context)

	profileToken := getProfileToken(context)
	position := getPosition(context)
	speed := getSpeed(context)

	reply, err := service.AbsoluteMove(&onvif_wsdl2go.AbsoluteMove{
		XMLName:      xml.Name{},
		ProfileToken: &profileToken,
		Position:     &position,
		Speed:        &speed,
	})

	if err != nil {
		//TODO error parse
		context.IndentedJSON(http.StatusInternalServerError, errorsProcessors.BuildError("000", err))
		return
	}

	log.Println(reply)

	//TODO reply body parse
	context.IndentedJSON(http.StatusOK, "{}")
}

func GetConfigurations(context *gin.Context) {
	service := getService(context)

	reply, err := service.GetConfigurations(&onvif_wsdl2go.GetConfigurations{})

	if err != nil {
		//TODO error parse
		context.IndentedJSON(http.StatusInternalServerError, errorsProcessors.BuildError("000", err))
		return
	}

	log.Println(reply)

	//TODO reply body parse
	context.IndentedJSON(http.StatusOK, "{}")
}

func GetPresets(context *gin.Context) {

	service := getService(context)

	profileToken := getProfileToken(context)

	reply, err := service.GetPresets(&onvif_wsdl2go.GetPresets{
		XMLName:      xml.Name{},
		ProfileToken: &profileToken,
	})

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, errorsProcessors.BuildError("000", err))
		return
	}

	log.Println(reply)

	//TODO reply body parse
	context.IndentedJSON(http.StatusOK, "{}")
}

func GotoHomePosition(context *gin.Context) {
	service := getService(context)

	profileToken := getProfileToken(context)
	speed := getSpeed(context)

	var reply, err = service.GotoHomePosition(&onvif_wsdl2go.GotoHomePosition{
		XMLName:      xml.Name{},
		ProfileToken: &profileToken,
		Speed:        &speed,
	})

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, errorsProcessors.BuildError("000", err))
		return
	}

	log.Println(reply)

	//TODO reply body parse
	context.IndentedJSON(http.StatusOK, "{}")
}

func GotoPreset(context *gin.Context) {

	service := getService(context)

	profileToken := getProfileToken(context)
	presetToken := getPresetToken(context)
	speed := getSpeed(context)

	reply, err := service.GotoPreset(&onvif_wsdl2go.GotoPreset{
		XMLName:      xml.Name{},
		ProfileToken: &profileToken,
		PresetToken:  &presetToken,
		Speed:        &speed,
	})

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, errorsProcessors.BuildError("000", err))
		return
	}

	log.Println(reply)

	//TODO reply body parse
	context.IndentedJSON(http.StatusOK, "{}")
}

// GetPresetTours may be not implemented on CAM
func GetPresetTours(context *gin.Context) {

	service := getService(context)

	profileToken := getProfileToken(context)

	reply, err := service.GetPresetTours(&onvif_wsdl2go.GetPresetTours{
		XMLName:      xml.Name{},
		ProfileToken: &profileToken,
	})

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, errorsProcessors.BuildError("000", err))
		return
	}

	log.Println(reply)

	//TODO reply body parse
	context.IndentedJSON(http.StatusOK, "{}")
}
