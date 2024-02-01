package ptz_request_processors

import (
	"github.com/gin-gonic/gin"
	"github.com/hooklift/gowsdl/soap"
	"mattue/onif-rest-api-go/onvif_wsdl2go"
	"os"
	"strconv"
	"time"
)

func getService(context *gin.Context) onvif_wsdl2go.PTZ {

	//TODO cam to url mapping
	//client := soap.NewClient(camUrl,
	client := soap.NewClient(os.Getenv("API_URL"),
		soap.WithTimeout(time.Second*5),
	)

	service := onvif_wsdl2go.NewPTZ(client)

	return service
}

func getPresetToken(context *gin.Context) (presetToken onvif_wsdl2go.ReferenceToken) {
	//TODO if preset-token is not provided
	presetToken = onvif_wsdl2go.ReferenceToken(context.Query("preset-token"))
	return
}

func getProfileToken(context *gin.Context) (profileToken onvif_wsdl2go.ReferenceToken) {
	//TODO if profile-token is not provided
	profileToken = onvif_wsdl2go.ReferenceToken(context.Query("profile-token"))
	return
}

func getPosition(context *gin.Context) (position onvif_wsdl2go.PTZVector) {

	panTilt := getPanTilt(context, "position")
	zoom := getZoom(context, "position")

	position = onvif_wsdl2go.PTZVector{
		PanTilt: &panTilt,
		Zoom:    &zoom,
	}

	return
}

func getPanTilt(context *gin.Context, prefix string) (vector onvif_wsdl2go.Vector2D) {

	x, err := strconv.ParseFloat(context.Query(prefix+"_x"), 32)

	if err != nil {
		//TODO
	}

	y, err := strconv.ParseFloat(context.Query(prefix+"_y"), 32)

	if err != nil {
		//TODO
	}

	vector = onvif_wsdl2go.Vector2D{
		X:     float32(x),
		Y:     float32(y),
		Space: "",
	}

	return
}

func getZoom(context *gin.Context, prefix string) (zoom onvif_wsdl2go.Vector1D) {
	zoomValue, err := strconv.ParseFloat(context.Query(prefix+"_zoom"), 32)

	if err != nil {
		//TODO
	}

	zoom = onvif_wsdl2go.Vector1D{
		X:     float32(zoomValue),
		Space: "",
	}

	return
}

func getSpeed(context *gin.Context) (speed onvif_wsdl2go.PTZSpeed) {

	panTilt := getPanTilt(context, "speed")
	zoom := getZoom(context, "speed")

	speed = onvif_wsdl2go.PTZSpeed{
		PanTilt: &panTilt,
		Zoom:    &zoom,
	}

	return
}
