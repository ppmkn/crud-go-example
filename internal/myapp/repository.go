package myapp
import (
	"database/sql"
	_"github.com/lib/pq"
)

func CreateUser(db *sql.DB, name, email string) error {
	query := "INSERT INTO users(name, email) VALUES($1, $2)"
	_, err := db.Exec(query, name, email)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(db *sql.DB, id int) (*User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	row := db.QueryRow(query, id)

	user := &User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(db *sql.DB, id int, name, email string) error {
	query := "UPDATE users SET name = $1, email = $2 WHERE id = $3"
	_, err := db.Exec(query, name, email, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(db *sql.DB, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}