package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Shc 定义核心流水表
type Shc struct {
	// define shc log for ist
	id       int
	msgtype  int
	pan      string
	amount   float32
	trandate string
	trace    int
	acquirer string
	issuer   string
	termid   string
	respcode int
}

func main() {
	db, err := sql.Open("mysql", "oasisadm:oasisadm123@/shclogdb?charset=utf8")
	checkErr(err)

	//rows, err := db.Query("SELECT * FROM shclog")
	rows, err := db.Query("SELECT msgtype,pan,trace FROM shclog")
	checkErr(err)

	for rows.Next() {
		var shc Shc

		//err = rows.Scan(&shc.id, &msgtype, &pan, &amount, &trandate, &trace, &acquirer, &issuer, &termid, &respcode)
		err = rows.Scan(&shc.msgtype, &shc.pan, &shc.trace)

		fmt.Println("row : ", shc.pan, shc.trace)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/*
  //插入数据
    stmt, err := db.Prepare("INSERT shclog SET pan=?,trace=?,msgtype=?")
    checkErr(err)

    res, err := stmt.Exec("62123456", "2123456", "210")
    checkErr(err)

    id, err := res.LastInsertId()
    checkErr(err)

    fmt.Println(id)

    //更新数据
    stmt, err = db.Prepare("update shclog set trace=? where uid=?")
    checkErr(err)

    res, err = stmt.Exec("62123456", id)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

*/
