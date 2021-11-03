package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	pb "dir/usermgmt"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func set(w http.ResponseWriter, req *http.Request) {
	var url string = req.URL.String()

	//json.Unmarshal(str, &params)
	values := strings.Split(url[5:], ",") //Get the data string,start where the data begin and split it to two values
	connectToServer(w, values[0]+" "+values[1])

}
func connectToServer(w http.ResponseWriter, name string) {

	//go func() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		fmt.Fprintln(w, "did not connect to server")
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &pb.Msg{Tkn: name} //connect the server
	resStream, err := c.CreateMsg(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling CreateMsg RPC: %v", err)
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
		log.Printf("Response from CreateMsg: %v", msg.GetStr())
		fmt.Fprintln(w, msg.GetStr())
	}
	//}()
}
func main() {
	http.HandleFunc("/set", set)

	http.ListenAndServe(":50050", nil)
}
