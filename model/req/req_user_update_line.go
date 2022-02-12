package req

type ReqUpdateUserLine struct {
	UserId         string `json:"id"`
	Avatar         string `json:"avatar"`
	FullName       string `json:"fullName"`
	Age            int    `json:"age"`
	Sex            string `json:"sex"`
	Height         int    `json:"height"`
	Job            string `json:"job"`
	Income         string `json:"income"`
	Marriage       string `json:"marriage"`
	Children       string `json:"children"`
	Home           string `json:"home"`
	Zodiac         string `json:"zodiac"`
	Status         string `json:"status"`
	Formality      string `json:"formality"`
	LinkFb         string `json:"linkFb"`
	LinkIs         string `json:"linkIs"`
	ZlPhone        string `json:"ZlPhone"`
	Address        string `json:"address"`
	Target         string `json:"target"`
	About          string `json:"about"`
	CountFollower  int    `json:"countFollower"`
	CountFollowing int    `json:"countFollowing"`
	CountLike      int    `json:"countLike"`
}
