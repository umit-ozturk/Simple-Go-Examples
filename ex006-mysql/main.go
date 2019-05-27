package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	fmt.Println("MySQL Example with Go.")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")

	if err != nil{
		panic(err.Error())
	}

	defer db.Close()
	
	fmt.Println("Successfully Connected to Database.")

}