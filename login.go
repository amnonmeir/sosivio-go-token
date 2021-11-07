package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/lib/pq"

	cr "dir/rest_api/car"
	mc "dir/rest_api/connect_to_server"
	dt "dir/rest_api/data"
	pn "dir/rest_api/person"
	tkn "dir/rest_api/tokens"
)

var Users = make(map[string]string)

var dbUsers = map[string]dt.User{}
var DB *sql.DB

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signin", signin)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/read_persons", readPersons)
	http.HandleFunc("/read_person", readPerson)
	http.HandleFunc("/create_person", createPerson)
	http.HandleFunc("/update_person", updatePerson)
	http.HandleFunc("/delete_person", deletePerson)
	http.HandleFunc("/read_cars", readCars)
	http.HandleFunc("/read_car", readCar)
	http.HandleFunc("/create_car", createCar)
	http.HandleFunc("/update_car", updateCar)
	http.HandleFunc("/delete_car", deleteCar)
	http.ListenAndServe(":8888", nil)
}
func index(w http.ResponseWriter, req *http.Request) {

	var url string = req.URL.String()
	fmt.Println(url)
}

func signin(w http.ResponseWriter, req *http.Request) {

	var url string = req.URL.String()
	u := dt.SetUser(url)
	key := u.UserName + "," + u.Password
	fmt.Println(u.UserName + "," + u.Password + "," + u.First + "," + u.Last)
	_, ok := Users[key]
	fmt.Println(ok)
	if ok {
		if tkn.CheckTokenValidity(Users[key]) {
			fmt.Println("The token is still valid")
			mc.ConnectToServer(w, u.First+" "+u.Last)
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
		mc.ConnectToServer(w, u.First+" "+u.Last)
	}
}
func readPersons(w http.ResponseWriter, req *http.Request) {
	var url string = req.URL.String()
	values := strings.Split(url[14:], ",")
	if !check_token_validity(values[0] + "," + values[1]) {
		return
	}
	var p = pn.Persons{}
	p = p.Read(DB)
	for _, person := range p.Per {
		fmt.Fprintln(w, person.ID, person.First, person.Last, person.Telephone)
	}
}
func readPerson(w http.ResponseWriter, req *http.Request) {
	var url string = req.URL.String()
	values := strings.Split(url[13:], ",")
	if !check_token_validity(values[0] + "," + values[1]) {
		return
	}
	var p = pn.Person{
		ID: values[2],
	}
	p.ReadRow(DB)
	fmt.Fprintln(w, p.ID, p.First, p.Last, p.Telephone)

}
func createPerson(w http.ResponseWriter, req *http.Request) {
	var url string = req.URL.String()
	values := strings.Split(url[15:], ",")
	if !check_token_validity(values[0] + "," + values[1]) {
		return
	}
	var p = pn.Person{
		ID:        values[2],
		First:     values[3],
		Last:      values[4],
		Telephone: values[5],
	}
	p.Create(DB)
}
func updatePerson(w http.ResponseWriter, req *http.Request) {
	var url string = req.URL.String()
	values := strings.Split(url[15:], ",")
	if !check_token_validity(values[0] + "," + values[1]) {
		return
	}
	var p = pn.Person{
		ID:        values[2],
		First:     values[3],
		Last:      values[4],
		Telephone: values[5],
	}
	p.Update(DB)
}
func deletePerson(w http.ResponseWriter, req *http.Request) {
	var url string = req.URL.String()
	values := strings.Split(url[15:], ",")
	if !check_token_validity(values[0] + "," + values[1]) {
		return
	}
	var p = pn.Person{
		ID:        values[2],
		First:     values[3],
		Last:      values[4],
		Telephone: values[5],
	}
	p.Delete(DB)
}

func readCars(w http.ResponseWriter, req *http.Request) {
	var url string = req.URL.String()
	values := strings.Split(url[11:], ",")
	if !check_token_validity(values[0] + "," + values[1]) {
		return
	}
	var c = cr.Cars{}
	c = c.Read(DB)
	for _, car := range c.Cr {
		fmt.Fprintln(w, car.ID, car.Model, car.Color, car.Year)
	}
}
func readCar(w http.ResponseWriter, req *http.Request) {
	var url string = req.URL.String()
	values := strings.Split(url[10:], ",")
	if !check_token_validity(values[0] + "," + values[1]) {
		return
	}
	var c = cr.Car{
		ID: values[2],
	}
	c.ReadRow(DB)
	fmt.Fprintln(w, c.ID, c.Model, c.Color, c.Year)

}
func createCar(w http.ResponseWriter, req *http.Request) {
	var url string = req.URL.String()
	values := strings.Split(url[12:], ",")
	if !check_token_validity(values[0] + "," + values[1]) {
		return
	}
	var c = cr.Car{
		ID:    values[2],
		Model: values[3],
		Color: values[4],
		Year:  values[5],
	}
	c.Create(DB)
}
func updateCar(w http.ResponseWriter, req *http.Request) {
	var url string = req.URL.String()
	values := strings.Split(url[12:], ",")
	if !check_token_validity(values[0] + "," + values[1]) {
		return
	}
	var c = cr.Car{
		ID:    values[2],
		Model: values[3],
		Color: values[4],
		Year:  values[5],
	}
	c.Update(DB)
}
func deleteCar(w http.ResponseWriter, req *http.Request) {
	var url string = req.URL.String()
	values := strings.Split(url[12:], ",")
	if !check_token_validity(values[0] + "," + values[1]) {
		return
	}
	var c = cr.Car{
		ID:    values[2],
		Model: values[3],
		Color: values[4],
		Year:  values[5],
	}
	c.Delete(DB)
}
func check_token_validity(key string) bool {

	_, ok := Users[key]
	if ok {
		if tkn.CheckTokenValidity(Users[key]) {
			return true
		} else {
			delete(Users, key)
			delete(dbUsers, key)
			return false
		}
	} else {
		return false
	}
}
func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://amnon:1234@localhost/dbclient?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB")
}
