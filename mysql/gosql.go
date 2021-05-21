package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	age  int
	name string
}

func initDb() (err error) {

	//连接数据库
	dsn := "root:4545123df@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf(" dsn format err : %v\n", err)
		return err

	}
	// 尝试连接
	err = db.Ping()
	if err != nil {
		fmt.Printf(" dsn ping format err : %v\n", err)
		return err

	}
	db.SetMaxOpenConns(10)
	return nil

}

func queryRow() {
	sqlStr := "SELECT id, name, age from user where id=?"
	// var u1 user

	for i := 0; i < 11; i++ {
		fmt.Printf("%v\n", i)
		db.QueryRow(sqlStr, 2)

	}

	// rowObj := db.QueryRow(sqlStr, 2)

	// rowObj.Scan(&u1.id, &u1.name, &u1.age)
	// fmt.Printf("id:%d name:%s age:%d\n", u1.id, u1.name, u1.age)

}

func queryMultiRow() {
	sqlStr := "select id, name, age from user where id > ?"

	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}

	defer db.Close()

	// rowserr := db.Query(sqlStr,0).Scan(&u.id, &u.name,&u.age)
	for rows.Next() {
		var u user

		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)

	}

}

func insertRow() {

	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "王五", 38)
	if err != nil {
		fmt.Printf("inset failed, err:%v\n", err)
		return
	}
	thId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf(" get lastinsert id failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d\n", thId)

}

func queryUpdate() {

	sqlStr := "UPDATE user set age =? where id =?"
	ret, err := db.Exec(sqlStr, "39", 3)
	if err != nil {
		fmt.Printf("Update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected fialed,  err:%v\n", err)
		return

	}
	fmt.Printf("update success, RowsAffected rows: %d\n", n)

}

func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("init db failed: %v\n", err)
		return
	}
	fmt.Println("连接数据库成功!!")

	// queryRow()
	// queryMultiRow()
	// insertRow()
	queryUpdate()

	defer db.Close()
}
