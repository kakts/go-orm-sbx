package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kakts/go-orm-sbx/gorm/service"
)

// form, jsonでのボディの変数名をタグで指定する
type RegistUserForm struct {
	Name string `form: "name" json:"name" binding:"required"`
	Age  uint8  `form: "age" json:"age" binding:"required"`
}

type GetUserByNameForm struct {
	Name string `form: "name" json:"name" binding:"required"`
}

type UserController struct{}

func (uc UserController) GetUser(c *gin.Context) {
	// user serviceを呼び出す
	var s = service.Service{}
	user := s.GetUser()
	c.JSON(200, gin.H{
		"Name": user.Name,
		"Age":  user.Age,
	})
}

func (uc UserController) GetUserByName(c *gin.Context) {
	// Bodyからnameを取る
	var s = service.Service{}
	var hoge GetUserByNameForm
	err := c.ShouldBindJSON(&hoge)
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
	}

	fmt.Printf("[GetUserByName] Got Name from req body. Name: %s\n", hoge.Name)
	users := s.GetUserByName(hoge.Name)
	c.JSON(200, gin.H{
		"counts": len(users),
		"Users":  users,
	})
}

func (uc UserController) RegistUser(c *gin.Context) {
	// userを登録する
	// NameとAgeのチェック
	var s = service.Service{}
	var hoge RegistUserForm
	err := c.ShouldBindJSON(&hoge)
	if err != nil {
		panic(err)
	}
	fmt.Println(hoge.Name)
	success := s.RegistUser(hoge.Name, hoge.Age)
	if !success {
		c.JSON(400, gin.H{
			"msg": "error",
		})

	}
	c.JSON(200, gin.H{
		"success": success,
	})
}
