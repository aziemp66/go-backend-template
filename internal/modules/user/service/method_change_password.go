package user_service

import "context"

func (userService *userService) ChangePassword(ctx context.Context, id, oldPassword, newPassword string) (err error) {
	user, err := userService.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	if err := userService.passwordManager.CheckPasswordHash(oldPassword, user.Password); err != nil {
		return err
	}

	if err := userService.passwordManager.PasswordValidation(newPassword); err != nil {
		return err
	}

	hashedPassword, err := userService.passwordManager.HashPassword(newPassword)
	if err != nil {
		return err
	}

	err = userService.userRepository.ChangePassword(ctx, id, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}
