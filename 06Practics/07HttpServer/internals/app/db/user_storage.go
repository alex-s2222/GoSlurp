package db

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"httpServer/internals/app/models"

	"log"

)


type UserStorage struct {
	databasePool *pgxpool.Pool
}


func NewUsersStorage(pool *pgxpool.Pool) *UserStorage{
	storage := new(UserStorage)
	storage.databasePool = pool
	return storage
}

func (storage *UserStorage) GetUserList(nameFilter string) []models.User {
	query := "SELECT id, name, rank FROM users"
	args := make([]interface{}, 0)
	if nameFilter != ""{
		query += "WHERE name LIKE $1"
		args = append(args, fmt.Sprintf("%%%s%%", nameFilter))
	}
	
	var result []models.User
	err := pgxscan.Select(context.Background(), storage.databasePool, &result, query, args...)
	
	if err != nil{ 
		log.Println(err)
	}
	return result
}

func (storage *UserStorage) GetUserById(id int64) models.User {
	query := "SELECT id, name, rank FROM users WHERE id = $1"
	
	var result models.User

	err := pgxscan.Get(context.Background(), storage.databasePool, &result, query, id)

	if err != nil {
		log.Println(err)
	}
	return result
}