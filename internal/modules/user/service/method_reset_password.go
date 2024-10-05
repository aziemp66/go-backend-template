package user_service

import "context"

func (userService *userService) ResetPassword(ctx context.Context, token string, newPassword string) (err error) {
	claims, err := userService.jwtManager.VerifyAuthToken(token)
	if err != nil {
		return err
	}

	user, err := userService.userRepository.GetUserByID(ctx, claims.ID)
	if err != nil {
		return err
	}

	if err := userService.passwordManager.PasswordValidation(newPassword); err != nil {
		return err
	}

	hashedPassword, err := userService.passwordManager.HashPassword(newPassword)
	if err != nil {
		return err
	}

	err = userService.userRepository.ChangePassword(ctx, user.Email, hashedPassword)
	if err != nil {
		return nil
	}

	return nil
}
