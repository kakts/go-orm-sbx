package service

import (
	"fmt"

	"github.com/kakts/go-orm-sbx/gorm_app/models"
	"github.com/kakts/go-orm-sbx/gorm_app/repository"
	"gorm.io/gorm"
)

type Service struct{}

func (s Service) GetUser() models.Users {
	db := repository.GetDB()

	fmt.Println("[getUser] get a user data from Users table.")
	userData := models.Users{}
	db.First(&userData)

	fmt.Printf("username = %s\n", userData.Name)
	return userData
}

type ResultUserNameAge struct {
	Id   uint
	Name string
	Age  uint8
}

// nameに一致したユーザ一覧を取得
func (s Service) GetUserByName(name string) []models.Users {
	db := repository.GetDB()

	userData := []models.Users{}
	db.Where("name = ?", name).Find(&userData)

	return userData
}

// nameに一致したユーザを取得
func (s Service) GetUserById(id uint) *models.Users {
	db := repository.GetDB()

	userData := models.Users{
		Model: gorm.Model{ID: id},
	}
	result := db.Where("id = ?", id).Find(&userData)

	if result.RowsAffected == 0 {
		return nil
	}
	return &userData
}

func (s Service) GetUsers() []models.Users {
	db := repository.GetDB()

	fmt.Println("[getUser] get a user data from Users table.")
	userList := []models.Users{}
	db.Find(&userList)

	return userList
}

// ユーザ登録
func (s Service) RegistUser(name string, age uint8) bool {
	db := repository.GetDB()
	user := &models.Users{
		Name: name,
		Age:  age,
	}
	result := db.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Printf("[RegistUser] successCount: %d, Name: %s, Age: %d", result.RowsAffected, user.Name, user.Age)

	return true
}
