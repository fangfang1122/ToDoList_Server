package routers

import (
	"ToDoList_Go/api"
	v1 "ToDoList_Go/api/v1"
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

	r.POST("/api/login", api.Login)

	route := r.Group("/api", jwt.UserRequired())
	{
		// 用户
		route.GET("/user/info", api.GetUserInfo)
		route.POST("/user/info", api.UpdateUser)
		route.POST("/user/avatar", api.UploadUserAvatar)

		// 清单
		route.GET("/lists", v1.GetAllList)
		route.POST("/list", v1.AddList)
		route.POST("list/:id", v1.UpdateList)
		route.DELETE("/list/:id", v1.DeleteList)

		// 清单
		route.GET("/tasks", v1.GetTaskList)
		route.POST("/task", v1.AddTask)
		route.POST("task/:id", v1.UpdateTask)
		route.DELETE("/task/:id", v1.DeleteTask)
		route.POST("/task/:id/file", v1.UploadTaskFile)
		route.POST("/task/:id/photo", v1.UploadTaskPhoto)
		route.POST("/task/:id/finish", v1.FinishTask)
		route.POST("/task/:id/cancel", v1.CancelTask)
	}

	return r
}
