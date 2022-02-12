package repo_impl

import (
	"context"
	"database/sql"
	"lienquanMess/banana"
	"lienquanMess/db"
	"lienquanMess/log"
	"lienquanMess/model"
	"lienquanMess/repository"
	"time"
)

type NewsFeedImpl struct {
	sql *db.Sql
}

func NewsFeedRepo(sql *db.Sql) repository.NewsFeedRepo {
	return &NewsFeedImpl{
		sql: sql,
	}
}

func (n *NewsFeedImpl) SelectAllNewsFeed(context context.Context, start string, userId string) ([]model.Posts, error) {
	var newsfeed []model.Posts
	err := n.sql.Db.SelectContext(context, &newsfeed,
		"SELECT * FROM posts WHERE user_id != $1 offset $2 limit 20", userId, start)
	if err != nil {
		log.Error(err.Error())
		return newsfeed, nil
	}
	return newsfeed, nil
}

func (n *NewsFeedImpl) AddPosts(context context.Context, post model.Posts, like model.Likes) (model.Posts, error) {
	statement := `
		INSERT INTO posts(id, user_id, access_modifier,content, image_url, created_at)
		VALUES(:id, :user_id, :access_modifier, :content, :image_url, :created_at)
	`
	post.CreatedAt = time.Now()
	_, err := n.sql.Db.NamedExecContext(context, statement, post)
	if err != nil {
		log.Error(err.Error())
		return post, banana.AddPostFail
	}
	return post, nil
}

func (n *NewsFeedImpl) SelectUserById(context context.Context, userId string) (model.AllUser, error) {
	var user model.AllUser

	err := n.sql.Db.GetContext(context, &user,
		"SELECT * FROM  userinfo WHERE id = $1", userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, banana.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}
