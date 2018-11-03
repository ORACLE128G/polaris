package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(192.168.1.9:3306)/analysis")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select id, `Name` from elastic limit 20000")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var col1, col2 []byte
	for rows.Next() {
		err = rows.Scan(&col1, &col2)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(col1), string(col2))
	}
}
