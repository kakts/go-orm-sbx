/**
 * web server
 */
package server

import (
	"github.com/gin-gonic/gin"

	"github.com/kakts/go-orm-sbx/gorm/handlers"
)

func Init() {
	r := router()

	r.Run()
}

// routing
func router() *gin.Engine {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	// /usersのハンドラー登録
	u := r.Group("/users")
	{
		controller := handlers.UserController{}
		u.GET("", controller.GetUser)
		u.GET("/:id", controller.GetUserById)
		u.POST("register", controller.RegistUser)
		u.GET("findname", controller.GetUserByName)
	}

	return r
}
