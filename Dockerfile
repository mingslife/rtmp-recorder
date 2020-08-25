FROM golang:1.12 as build
WORKDIR /app/rtmp-recorder
COPY . .
RUN GOPROXY=https://goproxy.io CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o rtmp-recorder cmd/recorder/main.go && chmod +x rtmp-recorder

FROM docker:19.03-dind as run
WORKDIR /app/rtmp-recorder
COPY --from=build /app/rtmp-recorder/rtmp-recorder .
EXPOSE 5000
ENTRYPOINT [ "/app/rtmp-recorder/rtmp-recorder" ]
