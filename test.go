package main

import . "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"


func main(){
	var db , err = sql.Open("mysql", "root:gasidegg@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
        Println(err)
    }
    var num , id int64
    var openid = "s"
    row , _ := db.Query("SELECT num,id FROM yd_alldaylog where openid like %?%",openid)
    for row.Next(){
    	row.Scan(&num,&id)
		Println(num)
		Println(id)
    }
	
}

