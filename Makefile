rtmp-recorder-arm:
	GOOS=linux GOARCH=arm go build -o rtmp-recorder-arm cmd/recorder/main.go

image-worker:
	docker build -t m1n9.vip/rtmp-recorder-worker:latest ./worker
