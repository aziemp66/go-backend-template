package user_repository_postgres

const (
	createUserQuery = `
	INSERT INTO user (email, password, name, address) VALUES($1, $2, $3, $4) RETURNING id
	`
	getUserByID = `
	SELECT (email, password, name, address) FROM user WHERE id = $1
	`
	getUserByEmail = `
	SELECT (email, password, name, address) FROM user WHERE email = $1
	`
	updateUserQuery = `
	UPDATE user set name = $2, address = $3 WHERE user = $1
	`
	changePasswordQuery = `
	UPDATE user set password = $2 WHERE email = $1
	`
	deleteUserQuery = `
	DELETE FROM user WHERE id = $1
	`

	verifyUserQuery = `
	UPDATE user set is_verified = true WHERE email = $1
	`
)
