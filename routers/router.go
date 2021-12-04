package routers

import (
	"ToDoList_Go/api"
	"ToDoList_Go/middleware/jwt"
	"ToDoList_Go/pkg/setting"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.RunMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.StaticFS("/api/upload", http.Dir("upload"))
	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello! welcome to use ToDoList~ made by funcfang",
		})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "page not found",
		})
	})

	r.POST("/api/login",api.Login)

	route:=r.Group("/api",jwt.UserRequired())
	{
		route.GET("/info",api.UserInfo)
	}

	return r
}
