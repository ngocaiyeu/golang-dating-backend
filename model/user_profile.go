package model

import "time"

type UserProfile struct {
	UserId         string    `json:"id" db:"id,omitempty"`
	Avatar         string    `json:"avatar" db:"avatar_url,omitempty"`
	FullName       string    `json:"fullName" db:"full_name, omitempty"`
	Age            int       `json:"age" db:"age,omitempty"`
	Sex            string    `json:"sex" db:"sex,omitempty"'`
	Height         int       `json:"height" db:"height,omitempty"`
	Job            string    `json:"job" db:"job,omitempty"`
	Income         string    `json:"income" db:"income,omitempty"`
	Marriage       string    `json:"marriage" db:"marriage,omitempty"`
	Children       string    `json:"children" db:"children,omitempty"`
	Home           string    `json:"home" db:"home,omitempty"`
	Zodiac         string    `json:"zodiac" db:"zodiac,omitempty"`
	Status         string    `json:"status" db:"status,omitempty"`
	Formality      string    `json:"formality" db:"formality,omitempty"`
	LinkFb         string    `json:"linkFb" db:"link_fb,omitempty"`
	LinkIs         string    `json:"linkIs" db:"link_is,omitempty"`
	ZlPhone        string    `json:"zlPhone" db:"zl_phone,omitempty"`
	Address        string    `json:"address" db:"address,omitempty"`
	Target         string    `json:"target" db:"target,omitempty"`
	About          string    `json:"about" db:"about,omitempty"`
	CountFollower  int       `json:"countFollower" db:"count_follower"`
	CountFollowing int       `json:"countFollowing" db:"count_following"`
	CountLike      int       `json:"countLike" db:"count_like"`
	CreatedAt      time.Time `json:"-" db:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"-" db:"updated_at,omitempty"`
}
