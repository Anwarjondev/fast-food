package repository

import (
	"github.com/Anwarjondev/fast-food/internal/db"
)

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func CreateUser(email, password string) (int, error) {
	var id int
	err := db.DB.QueryRow(`Insert into users(email, password, is_active) values($1, $2, false) returning id`, email, password).Scan(&id)
	return id, err
}

func SaveToken(userId int, token string) error {
	_, err := db.DB.Exec(`Update users set token = $1 where id = $2`, token, userId)
	return err
}

func GetUserByToken(token string) (*User, error) {
	var user User
	err := db.DB.Get(&user, `Select id, email from users where token = $1`, token)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func SaveConfirmaion(userId int, code int) error {
	_, err := db.DB.Exec(`Insert into confirm(user_id, code, is_passwed, created_at) values($1, $2, false, now())`, userId, code)
	return err
}

func ActiveUser(UserID int) error {
	_, err := db.DB.Exec(`Update "users" Set is_active = true where id = $1`, UserID)
	return err
}

func CheckCode(UserID, code int) (bool, error) {
	var count int
	err := db.DB.Get(&count, `Select count(*) from confirm where user_id = $1 and code = $2 and is_passwed = false`, UserID, code)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func MarkCodeAsPassed(UserID int) error {
	_, err := db.DB.Exec(`Update confirm set is_passwed = true where user_id = $1`, UserID)
	return err
}

func LoginUser(email, password string) (int, error) {
	var id int
	err := db.DB.Get(&id, `Select id from users where email = $1 and password = $2 and is_active = true`, email, password)
	if err != nil {
		return 0, err
	}
	_, err = db.DB.Exec(`Update "users" set is_logged_in = true where id = $1`, id)
	return id, err
}

func LogoutUser(UserID int) error {
	_, err := db.DB.Exec(`Update "users" set is_logged_in = false where id = $1`, UserID)
	return err
}

func UpdatePassword(userID int, newPassword string) error {
	_, err := db.DB.Exec(`Update users set password = $1 where id = $2`, newPassword, userID)
	return err
}

func GetUserIDByCode(code int) (int, error) {
	var userID int
	err := db.DB.Get(&userID, `
		SELECT user_id 
		FROM confirm 
		WHERE code = $1 
		AND is_passwed = false 
		AND created_at > NOW() - INTERVAL '1 minute'
	`, code)
	return userID, err
}
