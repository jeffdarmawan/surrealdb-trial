package main

import (
	"fmt"

	"github.com/surrealdb/surrealdb.go"

	"github.com/surrealdb/surrealdb.go/pkg/conn/gorilla"
	"github.com/surrealdb/surrealdb.go/pkg/marshal"
)

type User struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func main() {
	// init connection
	conn := gorilla.Create()

	db, err := surrealdb.New("ws://localhost:8000/rpc", conn)
	if err != nil {
		panic(err)
	}

	fmt.Println("passwed new")
	// signin db
	if _, err = db.Signin(&surrealdb.Auth{
		Username: "root",
		Password: "root",
	}); err != nil {
		panic(err)
	}

	fmt.Println("passwed signin")

	// use db ?
	if _, err := db.Use("test", "test"); err != nil {
		panic(err)
	}

	fmt.Println("passwed use")

	// User object
	user := User{
		Name:    "John",
		Surname: "Doe",
	}

	// insert user
	data, err := db.Create("user", user)
	if err != nil {
		panic(err)
	}

	// unmarshal data
	createdUser := make([]User, 1)
	err = marshal.Unmarshal(data, &createdUser)
	if err != nil {
		panic(err)
	}

	// get user by ID
	data, err = db.Select(createdUser[0].ID)
	if err != nil {
		panic(err)
	}

	// unmarshal data
	selectedUser := new(User)
	err = marshal.Unmarshal(data, &selectedUser)
	if err != nil {
		panic(err)
	}

	fmt.Println("passwed selectedUser: ", selectedUser)

	// update user
	changes := map[string]string{"name": "Jane"}
	if _, err = db.Merge(selectedUser.ID, changes); err != nil {
		panic(err)
	}

	if _, err = db.Query("SELECT * FROM $record", map[string]interface{}{
		"record": createdUser[0].ID,
	}); err != nil {
		panic(err)
	}
	fmt.Println("passwed query")

	if _, err = db.Delete(selectedUser.ID); err != nil {
		panic(err)
	}

}
