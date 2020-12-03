package user

type UserData struct {
	UserID    string `json:"user_id"`
	AuthToken string `json:"auth_token"`
	Name      string `json:"name"`
}
