package main

import (
	"net/http"
	"strconv"

	"rtmp-recorder/pkg/conf"
	"rtmp-recorder/pkg/models"
	"rtmp-recorder/pkg/router"
)

func main() {
	c := conf.ParseConfig()

	models.MasterId = c.MasterId

	router := router.NewRouter(c)

	port := strconv.Itoa(c.Port)
	http.ListenAndServe("0.0.0.0:"+port, router)
}
