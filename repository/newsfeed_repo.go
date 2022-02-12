package repository

import (
	"context"
	"lienquanMess/model"
)

type NewsFeedRepo interface {
	SelectAllNewsFeed(context context.Context, start string, userId string) ([]model.Posts, error)
	AddPosts(context context.Context, post model.Posts, like model.Likes) (model.Posts, error)
	SelectUserById(context context.Context, userId string) (model.AllUser, error)
}
