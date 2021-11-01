package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	pb "example.com/go-usermgmt-grpc"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

var params Response
var mySigningKey = []byte("ValidToken")

type Response struct {
	Name     string
	Password string
}

func GenerateJWT() (string, error) { //generate token and return it
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	fmt.Printf("Name: %s\n", params.Name)
	fmt.Printf("Password: %s\n", params.Password)
	claims["authorized"] = true
	claims["client"] = params.Name
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
func set(w http.ResponseWriter, req *http.Request) {
	params = Response{}
	var url string = req.URL.String()

	//json.Unmarshal(str, &params)
	values := strings.Split(url[5:], ",") //Get the data string,start where the data begin and split it to two values

	params.Name = values[0]
	params.Password = values[1]
	if checkUserValidity(w) {
		connectToServer(w)
	}

}
func checkUserValidity(w http.ResponseWriter) bool { //check if the u+p is valid
	if params.Name == "amnon" && params.Password == "meir" {
		return true
	}
	fmt.Fprintln(w, "User not valid")
	return false
}
func connectToServer(w http.ResponseWriter) { //communicate with the server to return values

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		fmt.Fprintln(w, "did not connect to server")
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
		fmt.Fprintln(w, "Failed to generate token")
		return
	}

	r, err := c.CreateMsg(ctx, &pb.Msg{Tkn: validToken}) //connect the server
	if err != nil {
		log.Fatalf("could not create user: %v", err)
		fmt.Fprintln(w, "could not create user")
		return
	}
	for _, str := range r.GetStr() { //return string array as an answer
		log.Printf(`token Details: %s`, str)
		fmt.Fprintln(w, str)
	}
}
func main() {
	http.HandleFunc("/set", set)

	http.ListenAndServe(":50050", nil)
}
