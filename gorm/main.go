package main

import (
	"fmt"

	"github.com/kakts/go-orm-sbx/gorm/models"
	"github.com/kakts/go-orm-sbx/gorm/repository"
	"gorm.io/gorm"
)

const (
	DBUser     = "user"
	DBPass     = "pass"
	DBProtocol = "tcp(127.0.0.1:3306)"
	DBName     = "test_database"
)

var users models.Users

// autoMigrate migrates DB tables from models
func autoMigrate(db *gorm.DB) {
	fmt.Println("[autoMigrate] migrate DB tables from models.")
	db.AutoMigrate(&users)
}

// addUsers add 2 users data with transaction
func addUsers(db *gorm.DB) {
	// Users2件の作成を1トランザクションで行う
	tx := db.Begin()
	user1 := models.Users{
		Name: "test1",
		Age:  8,
	}
	user2 := models.Users{
		Name: "test2",
		Age:  8,
	}
	tx.Create(&user1)
	tx.Create(&user2)

	success := true
	if success {
		// 正常時コミット成功
		fmt.Println("Succeeded. commit will be executed.")
		tx.Commit()
	} else {
		// ロールバック時の処理
		fmt.Println("Failed. rollback will be executed.")
		tx.Rollback()
	}
	fmt.Println("Transaction done!")
}

// getUser gets a data encounted first from Users table
func getUser(db *gorm.DB) {

	fmt.Println("[getUser] get a user data from Users table.")
	userData := models.Users{}
	db.First(&userData)

	fmt.Printf("username = %s\n", userData.Name)
}

// getUsersByName gets Data corresponding to a name
func getUsersByName(db *gorm.DB, name string) {

	usersData := []models.Users{}
	// 指定したnameの値に一致するデータを取得
	db.Where("name = ?", name).Find(&usersData)

	fmt.Printf("[getUsersByName] len = %d\n", len(usersData))

	for i := 0; i < len(usersData); i++ {
		userData := usersData[i]
		// TODO userData.IDへのアクセス方法
		fmt.Printf("users.name = %s, age = %d\n", userData.Name, userData.Age)
	}
}

func main() {
	d := repository.NewDB()
	db := d.Connect()

	// DBの自動マイグレート
	// modelから自動的にテーブルの作成・修正し、最新の状態に保つ
	autoMigrate(db)

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

	addUsers(db)

	getUser(db)

	getUsersByName(db, "test1")
}
