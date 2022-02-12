package model

import "time"

type User struct {
	UserId    string    `json:"id" db:"id,omitempty"`
	FullName  string    `json:"fullName,omitempty" db:"full_name,omitempty"`
	Phone     string    `json:"_,omitempty" db:"phone,omitempty"`
	Password  string    `json:"-" db:"password,omitempty"`
	Token     string    `json:"token,omitempty"`
	CreatedAt time.Time `json:"-" db:"created_at,omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at,omitempty"`
}
