package main

import (
	"context"
	"github.com/god-jay/gools/_examples/gin-protoc-project/pb"
	"log"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
)

func main() {
	callHTTP()
}
func callHTTP() {
	conn, err := transhttp.NewClient(
		context.Background(),
		transhttp.WithMiddleware(
			recovery.Recovery(),
		),
		transhttp.WithEndpoint("127.0.0.1:8001"),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewUserServiceHTTPClient(conn)
	reply, err := client.GetIndex(context.Background(), &pb.IndexRequest{Id: 999})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[http] SayHello: %s\n", reply.String())
}
