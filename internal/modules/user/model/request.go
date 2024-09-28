package user_model

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6"`
}

type UpdateUserRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6"`
}

type ChangePasswordRequest struct {
	Email       string `json:"email" binding:"required,email"`
	OldPassword string `json:"old_password" binding:"required,gte=6"`
	NewPassword string `json:"new_password" binding:"required,gte=6"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type GetUserRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type ResetPasswordRequest struct {
	Token string `uri:"token" binding:"required"`
}
