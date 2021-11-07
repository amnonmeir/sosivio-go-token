package makeConnection

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	pb "dir/usermgmt"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50050"
)

func ConnectToServer(w http.ResponseWriter, user string) {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &pb.ServerStreamingCall{Appeal: user} //connect the server
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
		fmt.Println("This is not finished yet")
	}
}
