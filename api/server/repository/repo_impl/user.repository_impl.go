package repoimpl

import (
	"bookz/api/server/repository"

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

func (u *UserRepositoryImpl) Login() {

}

func (u *UserRepositoryImpl) Register() {

}

func (u *UserRepositoryImpl) IsLogin() {

}
