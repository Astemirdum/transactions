package v1

type UserRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type User struct {
	ID       int    `db:"id"`
	Password string `db:"hash_password"`
	Email    string `db:"email"`
}
