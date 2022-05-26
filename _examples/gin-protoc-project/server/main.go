package main

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jay/gools/_examples/gin-protoc-project/pb"
	"log"
)

type server struct{}

func (s *server) GetIndex(c *gin.Context, request *pb.IndexRequest) (*pb.IndexResponse, error) {
	return &pb.IndexResponse{
		Id:    request.Id,
		Name:  "1",
		Email: "22",
		Phone: "333",
	}, nil
}

func main() {
	s := &server{}

	r := gin.New()

	pb.RegisterUserServiceHTTPServer(r, s)

	if err := r.Run(":8001"); err != nil {
		log.Fatal("failed to start gin server", err)
	}
}
