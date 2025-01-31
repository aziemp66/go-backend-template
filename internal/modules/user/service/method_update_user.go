package user_service

import (
	util_error "backend-template/util/error"
	"context"
	"database/sql"
	"fmt"
)

func (userService *userService) UpdateUser(ctx context.Context, id, name, address string) (err error) {
	_, err = userService.userRepository.GetUserByID(ctx, id)
	if err == sql.ErrNoRows {
		return util_error.NewNotFound(fmt.Errorf("user with the id of %s is not found", id), "User not found")
	} else if err != nil {
		return err
	}

	err = userService.userRepository.UpdateUser(ctx, id, name, address)
	if err != nil {
		return err
	}

	return nil
}
