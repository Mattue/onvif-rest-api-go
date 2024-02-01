# onvif-rest-api-go

The REST API (with notes) of [ONVIF protocol](https://www.onvif.org/onvif/ver20/util/operationIndex.html), with some architecture concessions for simple integration with [Nightbot](https://nightbot.tv/)

Under the hood [ONVIF protocol](https://www.onvif.org/onvif/ver20/util/operationIndex.html) is, I think I may say, a SOAP API.

Go definitions is generated by [gowsdl tool](https://github.com/hooklift/gowsdl).

# How to run

### 1. Clone and build image

```
docker build -t onvif-rest-api-go:latest .
```

### 2. Configure docker-compose

Make changes in `docker-compose.yml` file according to your needs

### 3. Run docker-compose

```
docker-compose up -d
```

# Implemented operations

Main focus is to provide transparent transition from SOAP to REST. So not many validations implemented.

## [PTZ](https://www.onvif.org/onvif/ver20/ptz/wsdl/ptz.wsdl)

### [AbsoluteMove](https://www.onvif.org/onvif/ver20/ptz/wsdl/ptz.wsdl#op.AbsoluteMove)

Should be POST with corresponding JSON body, but Nightbot can't do POST's

```
GET /v1/:camera/absolute-move?profile-token=xxx&
                                position-x=***&
                                position-y=***&
                                position-zoom=**&
                                speed-x=***&
                                speed-y=***&
                                speed-zoom=***
```

---

### [GetConfigurations](https://www.onvif.org/onvif/ver20/ptz/wsdl/ptz.wsdl#op.GetConfigurations)

```
GET /v1/:camera/get-configurations
```

---

### [GetPresets](https://www.onvif.org/onvif/ver20/ptz/wsdl/ptz.wsdl#op.GetPresets)

```
GET /v1/:camera/get-presets?profile-token=***
```

---

### GetPresetTours

```
GET /v1/:camera/get-preset-tours?profile-token=***
```

---

### GotoHomePosition

Should be POST with corresponding JSON body, but Nightbot can't do POST's

```
GET /v1/:camera/goto-home-position?profile-token=***&
                                    speed-x=***&
                                    speed-y=***&
                                    speed-zoom=***
```

---

### GotoPreset

Should be POST with corresponding JSON body, but Nightbot can't do POST's

```
GET /v1/:camera/goto-preset?profile-token=***&
                            preset-token=***&
                            speed-x=***&
                            speed-y=***&
                            speed-zoom=***
```

---

# Features

Multiple ONVIF endpoint support

To configure multiple endpoint support provide environment variables in format
```
cam1=https://camera1.host/onvif/device_service
cam2=https://camera2.host/onvif/device_service
prod=https://camera.prod.host/onvif/device_service
```

After start your cameras will be available at 
```
https://service.host/v1/cam1/***
https://service.host/v1/cam2/***
https://service.host/v1/prod/***
```

# Missing

* Response body from ONVIF parsing
* Error descriptions and correct error codes
* HTTP error code transparency

# TODO

* error on wrong endpoint
* log backend requests
* forward backend http codes