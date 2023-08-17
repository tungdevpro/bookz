package routes

import (
	"bookz/api/server/handler"
	repoImpl "bookz/api/server/repository/repo_impl"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRouter(r *gin.Engine, group *gin.RouterGroup, db *mongo.Client) {
	route := group.Group("/user")

	userHandler := handler.UserHandler{
		UserRepository: repoImpl.NewUserRepository(db),
	}

	fmt.Println("repoImpl::: ", userHandler.UserRepository)

	{
		route.POST("/login", userHandler.Login)
		route.POST("/register", userHandler.Register)
	}
}
