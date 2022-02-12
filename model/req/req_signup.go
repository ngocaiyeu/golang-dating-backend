package req

type ReqSignUp struct {
	FullName string `json:"fullName" validate:"required"`
	Phone    string `json:"phone,omitempty" validate:"required,e164"`
	Password string `json:"password,omitempty" validate:"pwd"`
}
