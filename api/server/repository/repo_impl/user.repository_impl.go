package repoimpl

import (
	"bookz/api/server/models"
	"bookz/api/server/repository"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) repository.UserRepository {
	return &UserRepositoryImpl{
		client: client,
	}
}

func (repo *UserRepositoryImpl) Login(ctx *gin.Context) {

}

func (repo *UserRepositoryImpl) Register(ctx *gin.Context, model models.User) {
	col := repo.client.Database("bookz").Collection("users")
	filter := bson.D{{Key: "email", Value: model.Email}}

	var user models.User

	col.InsertOne(ctx, model)

	if err := col.FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("record does not exist")
			return
		}
		fmt.Println("err: ", err)
	}
}

func (repo *UserRepositoryImpl) IsLogin(ctx *gin.Context) {

}
