package user_repository

import (
	"context"

	user_model "backend-template/internal/modules/user/model"
)

// UserRepository defines the set of operations for interacting with the user data at the repository level.
// This interface is typically implemented by a data layer that interacts with a database or external storage.
type UserRepository interface {

	// CreateUser creates a new user in the repository with the given details.
	// Returns the newly created user's ID or an error if the operation fails.
	CreateUser(ctx context.Context, email, hashedPassword, name, address string) (id string, err error)

	// GetUserByID retrieves a user by their unique ID.
	// Returns the user details wrapped in a user_model.User struct, or an error if not found.
	GetUserByID(ctx context.Context, id string) (res user_model.User, err error)

	// GetUserByEmail retrieves a user by their email address.
	// Returns the user details wrapped in a user_model.User struct, or an error if not found.
	GetUserByEmail(ctx context.Context, email string) (res user_model.User, err error)

	// ChangePassword updates a user's password in the repository.
	// Takes the user's email and the new hashed password, returning an error if the update fails.
	ChangePassword(ctx context.Context, id, hashedPassword string) (err error)

	// UpdateUser updates the user's name and address information.
	// Takes the user's ID and the updated name and address values. Returns an error if the update fails.
	UpdateUser(ctx context.Context, id string, name, address string) (err error)

	// DeleteUser removes a user from the repository by their ID.
	// Returns an error if the deletion operation fails.
	DeleteUser(ctx context.Context, id string) (err error)

	// VerifyUser marks a user as verified based on their email address.
	// Typically used after a successful email verification process.
	VerifyUser(ctx context.Context, id string) (err error)
}
