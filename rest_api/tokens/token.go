package token

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("ValidToken")

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
func CheckTokenValidity(tkn string) bool {

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
