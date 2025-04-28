package processors

import (
	"errors"
	"httpServer/internals/app/db"
	"httpServer/internals/app/models"
)

type UserProcessor struct {
	storage *db.UserStorage
}

func NewUserProcessor(storage *db.UserStorage) *UserProcessor{
	processor := new(UserProcessor)
	processor.storage = storage
	return processor
}

func (processor *UserProcessor) CreateUser(user models.User) (models.User, error) {
	if user.Name == ""{
		return user, errors.New("name should not be empty")
	}

	return processor.storage.GetUserById(3), nil
}