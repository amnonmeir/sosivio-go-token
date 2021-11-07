package person

import (
	"database/sql"
	"fmt"
)

type Data_base_commands interface{}
type Data_base interface {
	Read(DB *sql.DB)
	ReadRow(DB *sql.DB)
	Create(DB *sql.DB)
	Update(DB *sql.DB)
	Delete(DB *sql.DB)
}
type Person struct {
	ID        string `Json: "id"`
	First     string `Json: "first"`
	Last      string `Json: "last"`
	Telephone string `Json: "telephone"`
}
type Persons struct{ Per []Person }

func (p *Person) ReadRow(DB *sql.DB) {

	// Execute the query
	result := DB.QueryRow("SELECT id, first, last, telephone FROM T_person where id=" + p.ID)
	// for each row, scan the result into our tag composite object
	err := result.Scan(&p.ID, &p.First, &p.Last, &p.Telephone)
	if err != nil {
		panic(err.Error())
	}
}
func (p Persons) Read(DB *sql.DB) Persons {

	// Execute the query
	results, err := DB.Query("SELECT id, first, last, telephone FROM T_person")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	persons := make([]Person, 0)
	for results.Next() {
		person := Person{}
		// for each row, scan the result into our tag composite object
		err = results.Scan(&person.ID, &person.First, &person.Last, &person.Telephone)
		if err != nil {
			panic(err.Error())
		}
		persons = append(persons, person)
	}
	p.Per = persons
	return p
}
func (p Person) Create(DB *sql.DB) {
	s := "INSERT INTO T_person (id, first, last, telephone) VALUES('" + p.ID + "', '" + p.First +
		"', '" + p.Last + "', '" + p.Telephone + "')"
	_, err := DB.Query(s)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
func (p Person) Update(DB *sql.DB) {
	s := "UPDATE T_person SET first = '" + p.First + "', last = '" + p.Last +
		"', telephone = '" + p.Telephone + "' WHERE id = '" + p.ID + "'"
	fmt.Println(s)
	_, err := DB.Query(s)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
func (p Person) Delete(DB *sql.DB) {
	s := "DELETE FROM T_person WHERE id = '" + p.ID + "'"
	fmt.Println(s)
	_, err := DB.Query(s)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
