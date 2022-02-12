package model

import "time"

type Posts struct {
	Id             string    `json:"id,omitempty" db:"id"`
	UserId         string    `json:"userId,omitempty" db:"user_id"`
	AccessModifier string    `json:"accessModifier,omitempty" db:"access_modifier"`
	Content        string    `json:"content,omitempty" db:"content"`
	ImageUrl       string    `json:"imageUrl,omitempty" db:"image_url"`
	LikeCount      int       `json:"likeCount" db:"like_count"`
	CommentCount   int       `json:"commentCount" db:"comment_count"`
	CreatedAt      time.Time `json:"time" db:"created_at,omitempty"`
}
