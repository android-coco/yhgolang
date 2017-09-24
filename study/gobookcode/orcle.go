package main

import (
	"database/sql"
	// "fmt"
	_ "github.com/mattn/go-oci8"
	"log"
	// "os"
)

// func getDSN() string {
// 	var dsn string
// 	if len(os.Args) > 1 {
// 		dsn = os.Args[1]
// 		if dsn != "" {
// 			return dsn
// 		}
// 	}
// 	dsn = os.Getenv("GO_OCI8_CONNECT_STRING")
// 	if dsn != "" {
// 		return dsn
// 	}
// 	fmt.Fprintln(os.Stderr, `Please specifiy connection parameter in GO_OCI8_CONNECT_STRING environment variable,
// or as the first argument! (The format is user/name@host:port/sid)`)
// 	return "sys/123456@ORCL"
// }
func main() {
	// 为log添加短文件名,方便查看行数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("Oracle Driver example")

	//os.Setenv("NLS_LANG", "")

	// 用户名/密码@实例名  跟sqlplus的conn命令类似
	db, err := sql.Open("oci8", "youhao/youhao@39.108.218.80:1521/ORCL")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select 3.14, 'foo' from dual")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for rows.Next() {
		var f1 float64
		var f2 string
		rows.Scan(&f1, &f2)
		log.Println(f1, f2) // 3.14 foo
	}
	rows.Close()

	// 先删表,再建表
	db.Exec("drop table sdata")
	db.Exec("create table sdata(name varchar2(256))")

	db.Exec("insert into sdata values('中文')")
	db.Exec("insert into sdata values('1234567890ABCabc!@#$%^&*()_+')")

	rows, err = db.Query("select * from sdata")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var name string
		rows.Scan(&name)
		log.Printf("Name = %s, len=%d", name, len(name))
	}
	rows.Close()
}
