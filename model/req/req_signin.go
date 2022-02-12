package req

type ReqSignIp struct {
	Phone    string `json:"phone,omitempty" validate:"required,e164"`
	Password string `json:"password,omitempty" validate:"pwd"`
}
