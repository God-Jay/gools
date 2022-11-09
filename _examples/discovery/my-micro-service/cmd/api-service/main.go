package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	apiservice "github.com/god-jay/gools/_examples/discovery/my-micro-service/internal/service/api"
	"github.com/god-jay/gools/_examples/discovery/my-micro-service/pb"
	"github.com/god-jay/gools/discovery"
	"net/http"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	conf, err := discovery.ResolveConf(*configFile)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	srv := &http.Server{
		Addr:    conf.Service.ListenOn,
		Handler: router,
	}

	pb.RegisterApiUserServiceHTTPServer(router, apiservice.NewUserService(conf))

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
