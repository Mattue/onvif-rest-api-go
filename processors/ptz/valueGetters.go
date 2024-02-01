package ptz

import (
	"github.com/gin-gonic/gin"
	"github.com/hooklift/gowsdl/soap"
	"mattue/onif-rest-api-go/onvifGoDefinitions"
	"os"
	"strconv"
	"time"
)

func getService(context *gin.Context) onvifGoDefinitions.PTZ {

	//TODO check if os env does not exist
	client := soap.NewClient(os.Getenv(context.Param("camera")),
		soap.WithTimeout(time.Second*5),
	)

	service := onvifGoDefinitions.NewPTZ(client)

	return service
}

func getPresetToken(context *gin.Context) (presetToken onvifGoDefinitions.ReferenceToken) {
	//TODO if preset-token is not provided
	presetToken = onvifGoDefinitions.ReferenceToken(context.Query("preset-token"))
	return
}

func getProfileToken(context *gin.Context) (profileToken onvifGoDefinitions.ReferenceToken) {
	//TODO if profile-token is not provided
	profileToken = onvifGoDefinitions.ReferenceToken(context.Query("profile-token"))
	return
}

func getPosition(context *gin.Context) (position onvifGoDefinitions.PTZVector) {

	panTilt := getPanTilt(context, "position")
	zoom := getZoom(context, "position")

	position = onvifGoDefinitions.PTZVector{
		PanTilt: &panTilt,
		Zoom:    &zoom,
	}

	return
}

func getPanTilt(context *gin.Context, prefix string) (vector onvifGoDefinitions.Vector2D) {

	x, err := strconv.ParseFloat(context.Query(prefix+"_x"), 32)

	if err != nil {
		//TODO
	}

	y, err := strconv.ParseFloat(context.Query(prefix+"_y"), 32)

	if err != nil {
		//TODO
	}

	vector = onvifGoDefinitions.Vector2D{
		X:     float32(x),
		Y:     float32(y),
		Space: "",
	}

	return
}

func getZoom(context *gin.Context, prefix string) (zoom onvifGoDefinitions.Vector1D) {
	zoomValue, err := strconv.ParseFloat(context.Query(prefix+"_zoom"), 32)

	if err != nil {
		//TODO
	}

	zoom = onvifGoDefinitions.Vector1D{
		X:     float32(zoomValue),
		Space: "",
	}

	return
}

func getSpeed(context *gin.Context) (speed onvifGoDefinitions.PTZSpeed) {

	panTilt := getPanTilt(context, "speed")
	zoom := getZoom(context, "speed")

	speed = onvifGoDefinitions.PTZSpeed{
		PanTilt: &panTilt,
		Zoom:    &zoom,
	}

	return
}
