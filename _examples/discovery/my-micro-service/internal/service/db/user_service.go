package dbservice

import (
	"context"
	"github.com/god-jay/gools/_examples/discovery/my-micro-service/internal/db"
	"github.com/god-jay/gools/_examples/discovery/my-micro-service/pb"
	"log"
)

type UserService struct {
	db *db.Client
	pb.UnimplementedDBUserServiceServer
}

func NewUserService(db *db.Client) *UserService {
	return &UserService{db: db}
}

func (u *UserService) GetUserById(ctx context.Context, req *pb.DBUserRequest) (*pb.DBUserResponse, error) {
	log.Println("query ....")
	user, err := u.db.User(ctx).GetById(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DBUserResponse{
		Name: user.Name,
		Age:  user.Age,
		Sex:  user.Sex,
	}, nil
}
