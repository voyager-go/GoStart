package response

type UserMemberShowRes struct {
	Id           string
	Passport     string
	Email        string
	Status       string `json:"status"`
	StatusText   string `json:"status_text"`
	IfVerify     string `json:"if_verify"`
	IfVerifyText string `json:"if_verify_text"`
	Nickname     string `json:"nickname"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
