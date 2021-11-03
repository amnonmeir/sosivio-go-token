package main

import (
	"fmt"
	"net/http"

	dt "api/data_base"
	tkn "api/tokens"
)

var Users = make(map[string]string)

var dbUsers = map[string]dt.User{}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signin", signin)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, req *http.Request) {
}

func signin(w http.ResponseWriter, req *http.Request) {

	var url string = req.URL.String()
	u := dt.SetUser(url)
	key := u.UserName + "," + u.Password
	_, ok := Users[key]
	if ok {
		if tkn.CheckTokenValidity(Users[key]) {
			fmt.Println("The token is still valid")
			/////////////////////////////connectToServer(w, key)
		} else {
			fmt.Println("The token is not valid anymore")
			delete(Users, key)
			delete(dbUsers, key)
		}
	}
}

func signup(w http.ResponseWriter, req *http.Request) {

	var url string = req.URL.String()
	u := dt.SetUser(url)
	key := u.UserName + "," + u.Password
	_, ok := Users[key]
	if ok {
		if tkn.CheckTokenValidity(Users[key]) {
			fmt.Println("There is a user at the same name")
		}
	} else {
		fmt.Println("New token has been registered")
		validToken, err := tkn.GenerateJWT(key)
		if err != nil {
			fmt.Println("Failed to generate token")
			fmt.Fprintln(w, "Failed to generate token")
			return
		}
		dbUsers[u.UserName+","+u.Password] = u
		Users[key] = validToken
		///////////////////////////connectToServer(w, values[0])
	}
}
