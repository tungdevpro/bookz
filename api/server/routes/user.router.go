package routes

import (
	"bookz/api/server/handler"
	repoimpl "bookz/api/server/repository/repo_impl"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRouter(r *gin.Engine, group *gin.RouterGroup, db *mongo.Client) {
	route := group.Group("/user")

	userHandler := handler.UserHandler{
		UserRepository: repoimpl.NewUserRepository(db),
	}

	{
		route.POST("/login", userHandler.Login)
	}
}
