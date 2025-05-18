package models

type User struct {
	ID        int    `db:"id"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Is_Active bool   `db:"is_active"`
	Token     string `db:"token"`
}
type Confirm struct {
	ID        int  `db:"id"`
	Code      int  `db:"code"`
	User_id   int  `db:"user_id"`
	Is_Passed bool `db:"is_passed"`
}
