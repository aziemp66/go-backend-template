package user_model

type GetUserResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

type IDResponse struct {
	ID string `json:"id"`
}
