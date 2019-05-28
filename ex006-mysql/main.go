package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

type User struct {
	Name string
}

func main()  {
	fmt.Println("MySQL Example with Go.")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil{
		panic(err.Error())
	}

	defer db.Close()

	results, err := db.Query("SELECT name FROM users")
	if err != nil {
		panic(err.Error())
	}
	for results.Next(){
		var user User

		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.Name)
	}

	fmt.Println("Successfully Connected to Database.")

}