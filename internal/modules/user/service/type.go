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
	ChangePassword(ctx context.Context, id, oldPassword, newPassword string) error
}
