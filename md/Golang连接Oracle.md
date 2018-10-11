﻿# Golang连接Oracle

标签（空格分隔）： 学习

---

在此输入正文

## 一、安装MinGW
 http://blog.csdn.net/mecho/article/details/24305369
 我的是64位的
 安装我放在：
 C:\mingw-w64\mingw64
环境变量：

![环境变量][1]

## 二、安装OCI
http://www.oracle.com/technetwork/topics/winsoft-085727.html
![OCI版本][2]
 


## 三、配置go-oci8
 直接go get github.com/wendal/go-oci8（报错不用管），然后到go\src\github.com/wendal\go-oci8\windows下，将pkg-config.exe拷贝到C:\mingw-w64\mingw64\bin下，将oci8.pc复制到C:\mingw-w64\mingw64\lib\pkg-config\下，并且编辑oci8.pc：
```
# Package Information for pkg-config
prefix=修改为instantclient_11_1目录，如C:/androidtools/orcale/instantclient_12_2
exec_prefix=修改为instantclient_11_1目录，如C:/androidtools/orcale/instantclient_12_2
libdir=${exec_prefix}
includedir=${prefix}/sdk/include/

Name: OCI
Description: Oracle database engine
Version: 11.2
Libs: -L${libdir} -loci
Libs.private: 
Cflags: -I${includedir}
```
![pkg-config][3]


## 四、设置环境变量
环境变量path下添加C:/androidtools/orcale/instantclient_12_2和C:\mingw-w64\mingw64\bin的路径
添加PKG_CONFIG_PATH=C:\mingw-w64\mingw64\lib\pkg-config
添加TNS_ADMIN=C:/androidtools/orcale/instantclient_12_2\network\admin\

## 五、测试
go get github.com/mattn/go-oci8  执行不报错就对了 用这个驱动中文不乱码
```go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-oci8"
	"log"
	"os"
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

	// 用户名/密码@IP:端口/实例名  
	db, err := sql.Open("oci8", "song/123456@192.168.0.105:1521/ORCL")
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

```



# linux oci8

1,下载linux版OCI
http://www.oracle.com/technetwork/database/features/instant-client/index-097480.html

![此处输入图片的描述][4]


  解压到/usr/lib下
  ![此处输入图片的描述][5]


  [1]: http://note.youdao.com/yws/api/personal/file/83C4B66AE42744E5A7BDDC37D3AC054B?method=download&shareKey=13d92356f4a1796b00eaebe8c452da08
  [2]: http://note.youdao.com/yws/api/personal/file/412FC0462F43424989D9C161D20B4710?method=download&shareKey=d562877e203eb4f95eedcfedbef7dbe1
  [3]: http://note.youdao.com/yws/api/personal/file/E7B34677179846B19A324FB183DC491C?method=download&shareKey=412fb797b113f54eeb0f792a84aae0fb
  [4]: http://note.youdao.com/yws/api/personal/file/8D00DB9458924BE8A502AEF83FF93F13?method=download&shareKey=7dd714a0bcd10a38f187d69e390c0615
  [5]: http://note.youdao.com/yws/api/personal/file/68EBF5497B7749F7994016F1643B3EFB?method=download&shareKey=d4b490cdc15d6a55c0ca514f73383eb4
  
执行下面命令
```liunx
    ln /usr/lib/instantclient_12_2/libclntsh.so.12.1 /usr/lib/libclntsh.so
   ln /usr/lib/instantclient_12_2/libocci.so.12.1 /usr/lib/libocci.so
   ln /usr/lib/instantclient_12_2/libociei.so /usr/lib/libociei.so
  ln /usr/lib/instantclient_12_2/libnnz12.so /usr/lib/libnnz12.so
```
安装pkg-config
在 /usr/lib/pkg-config 目录下创建文件 oci8.pc，内容如下：
```
# Package Information for pkg-config

prefix=/usr/lib/instantclient_12_2
exec_prefix=/usr/lib/instantclient_12_2
libdir=${exec_prefix}
includedir=${prefix}/sdk/include/

Name: OCI
Description: Oracle database engine
Version: 12.2
Libs: -L${libdir} -lclntsh
Libs.private: 
Cflags: -I${includedir}
```
环境变量设置
vim /etc/profile
在最后加上
```
export ORACLE_HOME=/usr/lib/instantclient_12_2
export LD_LIBRARY_PATH=$ORACLE_HOME
export PKG_CONFIG_PATH=/usr/lib/pkg-config
```
source /etc/profile
执行驱动下载
```go
go get github.com/mattn/go-oci8
```


