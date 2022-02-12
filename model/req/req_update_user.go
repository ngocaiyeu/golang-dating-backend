package req

type ReqUpdateUser struct {
	FullName  string `json:"fullName,omitempty" validate:"required"`
	Sex       string `json:"sex,omitempty" validate:"required"`
	Age       int    `json:"age,omitempty" validate:"required"`
	Height    int    `json:"height,omitempty" validate:"required"`
	Job       string `json:"job,omitempty" validate:"required"`
	Income    string `json:"income,omitempty" validate:"required"`
	Marriage  string `json:"marriage,omitempty" validate:"required"`
	Children  string `json:"children,omitempty" validate:"required"`
	Home      string `json:"home,omitempty" validate:"required"`
	Zodiac    string `json:"zodiac,omitempty" validate:"required"`
	Status    string `json:"status,omitempty" validate:"required"`
	Formality string `json:"formality,omitempty" validate:"required"`
	LinkFb    string `json:"linkFb,omitempty"`
	LinkIs    string `json:"linkIs,omitempty"`
	Zalo      string `json:"zalo,omitempty"`
	Address   string `json:"address,omitempty" validate:"required"`
	Target    string `json:"target,omitempty" validate:"required"`
	About     string `json:"about,omitempty" validate:"required"`
	Desire    string `json:"desire,omitempty" validate:"required"`
}
