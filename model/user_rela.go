package model

type UserRela struct {
	Id        string   `json:"id,omitempty" db:"id"`
	Follower  []string `json:"Follower,omitempty" db:"follower"`
	Following []string `json:"Following,omitempty" db:"following"`
}
