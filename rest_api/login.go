package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

var users = make(map[string]string)
var dbUsers = map[string]user{}
var mySigningKey = []byte("ValidToken")

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func getUser(req *http.Request) user {
	var url string = req.URL.String()

	//json.Unmarshal(str, &params)
	values := strings.Split(url[5:], ",") //Get the data string,start where the data begin and split it to two values
	u := user{
		UserName: values[0],
		Password: values[1],
		First:    values[2],
		Last:     values[3],
	}
	dbUsers[u.UserName+","+u.Password] = u
	return u
}
func index(w http.ResponseWriter, req *http.Request) {
}

func bar(w http.ResponseWriter, req *http.Request) {

	u := getUser(req)
	key := u.UserName + "," + u.Password
	_, ok := users[key]
	if ok {
		if checkTokenValidity(users[key]) {
			fmt.Println("The token is still valid")
			/////////////connectToServer(w, key)
		} else {
			fmt.Println("The token is not valid anymore")
			delete(users, key)
			delete(dbUsers, key)
		}
	}
}

func signup(w http.ResponseWriter, req *http.Request) {

	u := getUser(req)
	key := u.UserName + "," + u.Password
	_, ok := users[key]
	if ok {
		if checkTokenValidity(users[key]) {
			fmt.Println("There is a user at the same name")
		}
	} else {
		fmt.Println("New token has been registered")
		validToken, err := GenerateJWT(key)
		if err != nil {
			fmt.Println("Failed to generate token")
			fmt.Fprintln(w, "Failed to generate token")
			return
		}
		users[key] = validToken
		//////////////////////////////////////////////////////////////////connectToServer(w, values[0])
	}
}
func GenerateJWT(name string) (string, error) { //generate token and return it
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	fmt.Printf("Name: %s\n", name)
	claims["authorized"] = true
	claims["client"] = name
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
func checkTokenValidity(tkn string) bool {

	fmt.Println("token:%v", tkn)
	token, _ := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) { //check token validity
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return mySigningKey, nil
	})
	if token.Valid { //on success
		return true
	} else { //on failure
		return false
	}
}
