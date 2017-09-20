package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lunny/godbc"
	"time"
)

func main() {
	conn, err := sql.Open("odbc", "driver={SQL Server};SERVER=127.0.0.1;UID=sa;PWD=youhao;DATABASE=test;")
	if err != nil {
		fmt.Println("Connecting Error")
		return
	} else {
		fmt.Println("Connecting Susses")
	}
	defer conn.Close()
	stmt, err := conn.Prepare("select * from [dbo].[user]")
	if err != nil {
		fmt.Println("Query Error", err)
		return
	}
	defer stmt.Close()
	row, err := stmt.Query()
	if err != nil {
		fmt.Println("Query Error", err)
		return
	}
	defer row.Close()
	for row.Next() {
		var name string
		var age int
		var datatime time.Time
		if err := row.Scan(&name, &age, &datatime); err == nil {
			fmt.Println(name)
			fmt.Println(age)
			fmt.Println(datatime)
		}
	}
	fmt.Printf("%s\n", "finish")
	return
}
