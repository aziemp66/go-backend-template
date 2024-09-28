package user_model

type GetUserResponse struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}
