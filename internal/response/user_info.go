package response

type UserInfoShowRes struct {
	Id        string
	Passport  string
	Nickname  string
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
