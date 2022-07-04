package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"
	pb "test.me/grpc_server/user"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	port        = 5050
)

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

type userInternalService struct {
	pb.UnimplementedUserInternalServiceServer
	savedUsers []*pb.User
}

func newServer() *userInternalService {
	s := &userInternalService{}
	return s
}

func (s *userInternalService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	for i := range s.savedUsers {
		if s.savedUsers[i].GetId() == in.GetUserId() {
			return &pb.GetUserResponse{
				Payload: &pb.UserResponsePayload{
					User: s.savedUsers[i],
				},
			}, nil
		}
	}

	return &pb.GetUserResponse{
		Payload: &pb.UserResponsePayload{
			User: &pb.User{
				Id:       "my nice id",
				Name:     "I don't exist",
				FullName: "I really don't exist",
				Aliases:  []string{"i", "a"},
			},
		},
	}, nil
}

func (s *userInternalService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &pb.User{
		Id:       RandStringBytes(20),
		Name:     in.Name,
		FullName: in.FullName,
		Aliases:  in.Aliases,
	}

	s.savedUsers = append(s.savedUsers, user)

	return &pb.CreateUserResponse{
		Payload: &pb.UserResponsePayload{
			User: user,
		},
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserInternalServiceServer(grpcServer, newServer())
	grpcServer.Serve(listener)
}
