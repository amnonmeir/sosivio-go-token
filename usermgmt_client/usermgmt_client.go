package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	pb "dir/usermgmt"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	port    = ":50050"
)

type UserManagementServer struct { //create struct of grpc
	pb.UnimplementedUserManagementServer
}

//get msg from the client
func (*UserManagementServer) ReturnMsgManyTimes(req *pb.ServerStreamingCall, stream pb.UserManagement_ReturnMsgManyTimesServer) error {

	name := req.GetAppeal()
	fmt.Println("name %v", name)
	connectToServer(name, stream)
	return nil
}

//connect to the sever
func connectToServer(name string, stream pb.UserManagement_ReturnMsgManyTimesServer) {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &pb.ServerStreamingCall{Appeal: name} //connect the server
	resStream, err := c.ReturnMsgManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling ReturnMsgManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		fmt.Println("msg %v", msg)
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from ReturnMsgManyTimes: %v", msg.GetAnswer())
		res := &pb.ServerStreamingReturnCall{Answer: msg.GetAnswer()}
		stream.Send(res) //send to the client
	}
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
