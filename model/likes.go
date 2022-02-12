package model

type Likes struct {
	Id     string   `json:"Id,omitempty" db:"id"`
	UserId []string `json:"UserId,omitempty" db:"user_id"`
	PostID string   `json:"PostID,omitempty" db:"post_id"`
}
