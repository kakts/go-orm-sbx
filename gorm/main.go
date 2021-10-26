package main

import (
	"fmt"

	"github.com/kakts/go-orm-sbx/gorm/repository"
)

const (
	DBUser     = "user"
	DBPass     = "pass"
	DBProtocol = "tcp(127.0.0.1:3306)"
	DBName     = "test_database"
)

func main() {
	d := repository.NewDB()
	db := d.Connect()

	/**
	 * 直接SQL文を書く
	 */
	rows, err := db.Raw("show tables").Rows()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			panic(err.Error())
		}

		fmt.Printf("tableName: %s\n", table)
	}

}
