package user_model

type User struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Address  string `db:"address"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
