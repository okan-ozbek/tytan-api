package user

import (
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) List() ([]*User, error) {
	rows, err := r.DB.Query("SELECT id, username, password, created_at FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user := &User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) Create(user *User) error {
	_, err := r.DB.Exec(
		"INSERT INTO user (username, password, created_at) VALUES (?, ?, ?)",
		user.Username,
		user.Password,
		user.CreatedAt,
	)

	return err
}

func (r *UserRepository) Read(id int) (*User, error) {
	row := r.DB.QueryRow("SELECT id, username, password, created_at FROM user WHERE id = ?", id)
	user := &User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Update(id int, user *User) error {
	_, err := r.DB.Exec(
		"UPDATE user SET username = ?, password = ?, created_at = ? WHERE id = ?",
		user.Username,
		user.Password,
		user.CreatedAt,
		id,
	)
	return err
}

func (r *UserRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM user WHERE id = ?", id)
	return err
}
