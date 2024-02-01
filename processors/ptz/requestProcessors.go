package ptz

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"log"
	"mattue/onif-rest-api-go/errorsProcessors"
	"mattue/onif-rest-api-go/onvifGoDefinitions"
	"net/http"
)

func AbsoluteMove(context *gin.Context) {
	service := getService(context)

	profileToken := getProfileToken(context)
	position := getPosition(context)
	speed := getSpeed(context)

	reply, err := service.AbsoluteMove(&onvifGoDefinitions.AbsoluteMove{
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

	reply, err := service.GetConfigurations(&onvifGoDefinitions.GetConfigurations{})

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

	reply, err := service.GetPresets(&onvifGoDefinitions.GetPresets{
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

	var reply, err = service.GotoHomePosition(&onvifGoDefinitions.GotoHomePosition{
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

	reply, err := service.GotoPreset(&onvifGoDefinitions.GotoPreset{
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

	reply, err := service.GetPresetTours(&onvifGoDefinitions.GetPresetTours{
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
