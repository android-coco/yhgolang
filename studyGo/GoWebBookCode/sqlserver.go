package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lunny/godbc"
)

func main() {
	conn, err := sql.Open("odbc", "driver={SQL Server};SERVER=127.0.0.1;UID=sa;PWD=youhao;DATABASE=test")
	if err != nil {
		fmt.Println("Connecting Error")
		return
	} else {
		fmt.Println("Connecting Susses")
	}
	defer conn.Close()
	stmt, err := conn.Prepare("select top 5 id from [dbo].[user]")
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
		var id int
		if err := row.Scan(&id); err == nil {
			fmt.Println(id)
		}
	}
	fmt.Printf("%s\n", "finish")
	return
}
