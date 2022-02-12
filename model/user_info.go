package model

type AllUser struct {
	UserId   string `json:"id" db:"id,omitempty"`
	Avatar   string `json:"avatar" db:"avatar_url,omitempty"`
	FullName string `json:"fullName" db:"full_name, omitempty"`
	Age      int8   `json:"age" db:"age,omitempty"`
	Sex      string `json:"sex" db:"sex,omitempty"'`
	Marriage string `json:"marriage" db:"marriage"`
	Address  string `json:"address" db:"address"`
}
