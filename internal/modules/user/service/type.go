package user_service

import (
	"context"

	user_model "backend-template/internal/modules/user/model"
)

// UserService defines the operations for managing users at the service level.
// This interface abstracts the higher-level business logic related to user management, authentication, and security.
type UserService interface {

	// Register registers a new user with the given details (email, password, name, address).
	// It hashes the password before saving and returns the new user's ID or an error if registration fails.
	Register(ctx context.Context, email, password, name, address string) (id string, err error)

	// GetUserByID fetches a user's details by their ID.
	// Returns a user response struct or an error if the user is not found.
	GetUserByID(ctx context.Context, id string) (res user_model.GetUserResponse, err error)

	// GetUserByEmail fetches a user's details by their email address.
	// Returns a user response struct or an error if the user is not found.
	GetUserByEmail(ctx context.Context, email string) (res user_model.GetUserResponse, err error)

	// Login authenticates a user by their email and password.
	// Returns a JWT token if authentication is successful or an error if login fails.
	Login(ctx context.Context, email, password string) (token string, err error)

	// UpdateUser updates a user's name and address information.
	// Returns an error if the update fails.
	UpdateUser(ctx context.Context, id, name, address string) error

	// ChangePassword allows a user to change their password by providing the old password and a new one.
	// Returns an error if the password change process fails (e.g., incorrect old password).
	ChangePassword(ctx context.Context, email, oldPassword, newPassword string) error

	// ForgotPassword initiates the forgot password process, typically sending a password reset link to the user's email.
	// Returns an error if the operation fails (e.g., email not found).
	ForgotPassword(ctx context.Context, email string) error

	// ResetPassword allows the user to reset their password using a valid reset token.
	// Returns an error if the token is invalid or the password reset fails.
	ResetPassword(ctx context.Context, token, newPassword string) (err error)

	// VerifyUser marks a user's email as verified, typically after they have confirmed their email address.
	// Returns an error if the verification process fails.
	VerifyUser(ctx context.Context, email string) (err error)
}
