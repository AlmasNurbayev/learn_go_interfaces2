package storage

import (
	"context"
	"interfaces2/internal/model"
	"log"

	"github.com/jmoiron/sqlx"
)

type UsersStorage struct {
	db *sqlx.DB
}

func NewUsersStorage(db *sqlx.DB) *UsersStorage {
	return &UsersStorage{db: db}
}

func (s *UsersStorage) Create(ctx context.Context, user model.User) (model.User, error) {
	log.Print("users created")
	return user, nil
}

func (s *UsersStorage) GetAll(ctx context.Context) ([]model.User, error) {
	mds := []model.User{
		{Id: 1, Name: "Name"},
		{Id: 2, Name: "Name2"},
	}
	log.Println("users GetAll", mds)
	return mds, nil
}

func (s *UsersStorage) GetById(ctx context.Context, id int64) (model.User, error) {
	md := model.User{
		Id: 1, Name: "Name",
	}
	log.Println("users get by id", md)
	return md, nil
}
