package data

import (
	"strings"
)

type User struct {
	UserName string
	Password string
	First    string
	Last     string
}

func SetUser(url string) User {

	//json.Unmarshal(str, &params)
	values := strings.Split(url[5:], ",") //Get the data string,start where the data begin and split it to two values
	u := User{
		UserName: values[0],
		Password: values[1],
		First:    values[2],
		Last:     values[3],
	}
	return u
}
