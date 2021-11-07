package person

import (
	"database/sql"
	"fmt"
)

type Car struct {
	ID    string `Json: "id"`
	Model string `Json: "first"`
	Color string `Json: "last"`
	Year  string `Json: "telephone"`
}
type Cars struct{ Cr []Car }

func (p *Car) ReadRow(DB *sql.DB) {

	// Execute the query
	result := DB.QueryRow("SELECT id, model, color, Year FROM T_car where id=" + p.ID)
	// for each row, scan the result into our tag composite object
	err := result.Scan(&p.ID, &p.Model, &p.Color, &p.Year)
	if err != nil {
		panic(err.Error())
	}
}
func (p Cars) Read(DB *sql.DB) Cars {

	// Execute the query
	results, err := DB.Query("SELECT id, model, color, Year FROM T_car")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	cars := make([]Car, 0)
	for results.Next() {
		car := Car{}
		// for each row, scan the result into our tag composite object
		err = results.Scan(&car.ID, &car.Model, &car.Color, &car.Year)
		if err != nil {
			panic(err.Error())
		}
		cars = append(cars, car)
	}
	p.Cr = cars
	return p
}
func (p Car) Create(DB *sql.DB) {
	s := "INSERT INTO T_car (id, model, color, year) VALUES('" + p.ID + "', '" + p.Model +
		"', '" + p.Color + "', '" + p.Year + "')"
	_, err := DB.Query(s)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
func (p Car) Update(DB *sql.DB) {
	s := "UPDATE T_car SET model = '" + p.Model + "', color = '" + p.Color +
		"', year = '" + p.Year + "' WHERE id = '" + p.ID + "'"
	fmt.Println(s)
	_, err := DB.Query(s)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
func (p Car) Delete(DB *sql.DB) {
	s := "DELETE FROM T_car WHERE id = '" + p.ID + "'"
	fmt.Println(s)
	_, err := DB.Query(s)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
