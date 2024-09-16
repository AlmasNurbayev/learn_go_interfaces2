package storage

import (
	"context"
	"interfaces2/internal/model"
	"log"

	"github.com/jmoiron/sqlx"
)

type ArticlesStorage struct {
	db *sqlx.DB
}

func NewArticlesStorage(db *sqlx.DB) *ArticlesStorage {
	return &ArticlesStorage{db: db}
}

func (s *ArticlesStorage) Create(ctx context.Context, article model.Article) (model.Article, error) {
	log.Print("article created")
	return article, nil
}

func (s *ArticlesStorage) GetAll(ctx context.Context) ([]model.Article, error) {
	mds := []model.Article{
		{Id: 1, Name: "Name", Title: "Title"},
		{Id: 2, Name: "Name2", Title: "Title2"},
	}
	log.Println("articles GetAll", mds)
	return mds, nil
}

func (s *ArticlesStorage) GetById(ctx context.Context, id int64) (model.Article, error) {
	md := model.Article{
		Id: 1, Name: "Name", Title: "Title",
	}
	log.Println("articles get by id", md)
	return md, nil
}
