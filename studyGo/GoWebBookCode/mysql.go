/*
* @Author: yhlyl
* @Date:   2017-09-09 22:38:09
* @Last Modified by:   yhlyl
* @Last Modified time: 2017-09-09 23:28:12
 */
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	//"time"
)

var (
	db  *sql.DB
	err error
	id  int64
)

//插入数据
func insert() {
	//返回准备要执行的sql操作,然后返回准备完毕的执行状态。
	//我们可以看到传入的参数都是=?对应的数据，这样做的方式可以在一定程度上防止SQL注入
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?, created=?")
	checkErr(err)
	//用来执行stmt准备好的SQL语句。
	res, err := stmt.Exec("youhao", "研发部门", "2017-09-09")

	checkErr(err)
	id, err = res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

//更新数据
func update() {
	//返回准备要执行的sql操作,然后返回准备完毕的执行状态。
	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	//用来执行stmt准备好的SQL语句。
	res, err := stmt.Exec("youhaoupdate", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	fmt.Println(affect)
}

//查询数据
func query() {
	//用来直接执行Sql返回Rows结果
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
}

//删除数据
func del() {
	//返回准备要执行的sql操作,然后返回准备完毕的执行状态。
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	//用来执行stmt准备好的SQL语句。
	res, err := stmt.Exec(id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()
}

func main() {
	/*
		user@unix(/path/to/socket)/dbname?charset=utf8
		user:password@tcp(localhost:5555)/dbname?charset=utf8
		user:password@/dbname
		user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
	*/
	db, err = sql.Open("mysql", "youhao:youhao@tcp(116.196.82.249:3306)/test?charset=utf8")
	checkErr(err)
	//插入数据
	insert()
	//更新数据
	update()
	//查询数据
	query()
	//删除数据
	del()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
