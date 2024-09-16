package parser

import (
	"context"
	"interfaces2/internal/model"
	"time"
)

type UserStorage interface {
	Create(ctx context.Context, user model.User) (model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
	GetById(ctx context.Context, id int64) (model.User, error)
}

type ArticleStorage interface {
	Create(ctx context.Context, article model.Article) (model.Article, error)
	GetAll(ctx context.Context) ([]model.Article, error)
	GetById(ctx context.Context, id int64) (model.Article, error)
}

type Parser struct {
	userStorage    UserStorage
	articleStorage ArticleStorage
	fetchInterval  time.Duration
}

func New(userStorage UserStorage, articleStorage ArticleStorage, fetchInterval time.Duration) *Parser {
	return &Parser{userStorage: userStorage, articleStorage: articleStorage, fetchInterval: fetchInterval}
}

func (f *Parser) Start(ctx context.Context) error {

	_, err := f.userStorage.Create(ctx, model.User{
		Name: "Name",
	})
	if err != nil {
		return err
	}

	_, err3 := f.userStorage.GetAll(ctx)
	if err3 != nil {
		return err3
	}

	_, err2 := f.userStorage.GetById(ctx, 1)
	if err2 != nil {
		return err2
	}

	_, err4 := f.articleStorage.Create(ctx, model.Article{
		Name:  "Name",
		Title: "Title",
	})
	if err4 != nil {
		return err4
	}

	_, err5 := f.articleStorage.GetAll(ctx)
	if err5 != nil {
		return err4
	}

	_, err6 := f.articleStorage.GetById(ctx, 1)
	if err6 != nil {
		return err4
	}

	return nil
}
