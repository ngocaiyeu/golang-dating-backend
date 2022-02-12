package req

type ReqAddPost struct {
	AccessModifier string `json:"AccessModifier" validate:"required"`
	Content        string `json:"Content,omitempty"`
	ImageUrl       string `json:"ImageUrl,omitempty"`
}
