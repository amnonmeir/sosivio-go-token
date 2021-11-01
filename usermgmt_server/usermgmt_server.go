package main

import (
	"context"
	"log"
	"net"

	jwt "github.com/dgrijalva/jwt-go"

	pb "example.com/go-usermgmt-grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var mySigningKey = []byte("ValidToken")

type UserManagementServer struct { //create struct of grpc
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateMsg(ctx context.Context, in *pb.Msg) (*pb.Returned, error) {
	var answer []string
	token, _ := jwt.Parse(in.GetTkn(), func(token *jwt.Token) (interface{}, error) { //check token validity
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return mySigningKey, nil
	})
	if token.Valid { //on success
		answer = []string{"Token valid", "Congratulations", "You logged the system", "Hope you are happy"}
		log.Println(answer)
	} else { //on failure
		answer = []string{"Token not valid"}
		log.Println(answer)
	}
	return &pb.Returned{Str: answer}, nil //return answer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //open listener connection to client
	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
