package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kakts/go-orm-sbx/gorm_app/models"
	"github.com/kakts/go-orm-sbx/gorm_app/service"
)

// form, jsonでのボディの変数名をタグで指定する
type RegistUserForm struct {
	Name string `form:"name" json:"name" binding:"required"`
	Age  uint8  `form:"age" json:"age" binding:"required"`
}

// GetUserByNameForm 名前をもとにユーザ情報を取得する用の型
type GetUserByNameForm struct {
	Name string `form:"name" json:"name" binding:"required"`
}

// GetUserByIdForm Idをもとにユーザ情報を取得する用の型
type GetUserByIdForm struct {
	Id uint `uri:"id" binding:"required"`
}

type UserController struct{}

// UserResponse Response用ユーザデータ
type UserResponse struct {
	Name string
	Age  uint
}

// UserByNameResponse Response用ユーザデータ
type UserByNameResponse struct {
	Counts uint
	Users  []models.Users
}

// @Summary ユーザ取得 GET:/users
// @Description 1件取得
// @Accept json
// @Produce json
// @Success 200 {object} UserResponse
// @Router /users [get]
// GET:/users 一件目のユーザを取得
func (uc UserController) GetUser(c *gin.Context) {
	// user serviceを呼び出す
	var s = service.Service{}
	user := s.GetUser()
	c.JSON(200, gin.H{
		"Name": user.Name,
		"Age":  user.Age,
	})
}

// @Summary GET:/users/findname 指定した名前に一致するユーザ一覧を取得
// @Description 指定した名前に一致するユーザ一覧を取得
// @Accept json
// @Produce json
// @Success 200 {object} UserByNameResponse
// @Router /users/findname [get]
func (uc UserController) GetUserByName(c *gin.Context) {
	// Bodyからnameを取る

	// TODO DI
	var s = service.Service{}
	var hoge GetUserByNameForm
	err := c.ShouldBindJSON(&hoge)
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
	}

	fmt.Printf("[GetUserByName] Got Name from req body. Name: %s\n", hoge.Name)
	users := s.GetUserByName(hoge.Name)
	c.JSON(200, gin.H{
		"Counts": len(users),
		"Users":  users,
	})
}

// TODO 指定したidのユーザデータのネガティブキャッシュ
// TODO 負荷試験 ネガティブキャッシュ有り版となし版での比較をする
// GET:/users/:id 指定したIDのユーザを取得
func (uc UserController) GetUserById(c *gin.Context) {
	// Bodyからnameを取る
	var s = service.Service{}
	var form GetUserByIdForm
	// uriからidをbindする
	err := c.ShouldBindUri(&form)
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	fmt.Printf("[GetUserById] Got Id from req uri. Id: %d\n", form.Id)
	user := s.GetUserById(form.Id)
	if user == nil {
		c.JSON(404, gin.H{
			"msg": fmt.Errorf("[GetUserById] status: 404 Cannot find user id:%d", form.Id).Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"Users": user,
	})
}

// POST:/users/register ユーザ登録
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
