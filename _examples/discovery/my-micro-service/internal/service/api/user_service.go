package apiservice

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/god-jay/gools/_examples/discovery/my-micro-service/pb"
	"github.com/god-jay/gools/discovery"
	"time"
)

type UserService struct {
	dbClient *discovery.Client
}

func NewUserService(conf *discovery.Conf) *UserService {
	dbClient := discovery.NewClient(conf, "DBService")
	return &UserService{dbClient: dbClient}
}

func (u *UserService) GetUser(c *gin.Context, request *pb.ApiUserRequest) (*pb.ApiUserResponse, error) {
	client := pb.NewDBUserServiceClient(u.dbClient)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(c, time.Second)
	defer cancel()
	r, err := client.GetUserById(ctx, &pb.DBUserRequest{
		Id: request.Id,
	})
	if err != nil {
		panic("could not greet: " + err.Error())
	}

	return &pb.ApiUserResponse{
		JwtToken: "abc",
		Name:     r.Name,
		Age:      r.Age,
		Sex:      r.Sex,
	}, nil
}
