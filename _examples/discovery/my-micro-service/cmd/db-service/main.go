package main

import (
	"flag"
	"github.com/god-jay/gools/_examples/discovery/my-micro-service/internal/db"
	dbservice "github.com/god-jay/gools/_examples/discovery/my-micro-service/internal/service/db"
	"github.com/god-jay/gools/_examples/discovery/my-micro-service/pb"
	"github.com/god-jay/gools/discovery"
	"google.golang.org/grpc"
	"log"
	"net"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")
var mysqlConf = flag.String("mf", "etc/mysql.yaml", "the mysql config file")

func main() {
	conf, err := discovery.ResolveConf(*configFile)
	if err != nil {
		panic(err)
	}
	discovery.RegService(conf)

	flag.Parse()
	lis, err := net.Listen("tcp", conf.Service.ListenOn)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	dbClient, err := db.NewClient(*mysqlConf)
	if err != nil {
		log.Fatalf("failed to connect to mysql: %v", err)
	}

	pb.RegisterDBUserServiceServer(s, dbservice.NewUserService(dbClient))

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
