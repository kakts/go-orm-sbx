/**
 * web server
 */
package server

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/kakts/go-orm-sbx/gorm_app/handlers"
)

//
func Init() {
	r := router()

	r.Run()
}

func router() *gin.Engine {

	r := gin.Default()

	// TODO 開発環境のみセットする
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

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
