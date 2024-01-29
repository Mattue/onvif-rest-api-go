FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
RUN go mod download
COPY . .
#RUN go build -o /onvif-rest-api-go
RUN go build -ldflags="-s -w" -o /app/onvif-rest-api-go main.go
#CMD ["sh"]

FROM alpine as prod
RUN apk update --no-cache && \
    apk add --no-cache tzdata
WORKDIR /app
COPY --from=builder /app/onvif-rest-api-go onvif-rest-api-go
#CMD [". /onvif-rest-api-go"]
#CMD ["sh"]
ENTRYPOINT ["./onvif-rest-api-go"]