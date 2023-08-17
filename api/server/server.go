package server

import (
	"bookz/api/server/routes"
	"bookz/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct{}

func (s *Server) Create() {

	db := &db.MongoDatabase{}
	db.Client = db.DBConnection()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "NotFound",
		})
	})

	v1 := r.Group("/v1")
	{
		routes.UserRouter(r, v1, db.Client)
	}

	r.Run(":3000")
}
