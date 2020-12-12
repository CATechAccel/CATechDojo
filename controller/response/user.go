package response

type UserResponse struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}

type GetUserResponse struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

type GetAllUserRespponse struct {
	Users []UserResponse `json:"users"`
}
