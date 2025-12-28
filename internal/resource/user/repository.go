package user

import (
	"database/sql"

	hashUtil "tytan-api/internal/util/hash"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *User) error {
	password, err := hashUtil.Hash(user.Password)
	if err != nil {
		return err
	}

	if _, err := r.DB.Exec(
		"INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)",
		user.Username,
		password,
		user.CreatedAt,
	); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(id int, user *User) error {
	password, err := hashUtil.Hash(user.Password)
	if err != nil {
		return err
	}

	if _, err := r.DB.Exec(
		"UPDATE users SET username = ?, password = ?, created_at = ? WHERE id = ?",
		user.Username,
		password,
		user.CreatedAt,
		id,
	); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

func (r *UserRepository) FindAll() ([]*User, error) {
	rows, err := r.DB.Query("SELECT id, username, password, created_at FROM users")
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

func (r *UserRepository) FindById(id int) (*User, error) {
	row := r.DB.QueryRow("SELECT id, username, password, created_at FROM users WHERE id = ?", id)

	user := &User{}
	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindByUsername(username string) (*User, error) {
	row := r.DB.QueryRow("SELECT * FROM users WHERE username = ?", username)

	user := &User{}
	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindByCredentials(username string, password string) (*User, error) {
	password, err := hashUtil.Hash(password)
	if err != nil {
		return nil, err
	}

	row := r.DB.QueryRow(
		"SELECT id, username, password, created_at FROM users WHERE username = ? AND password = ?",
		username,
		password,
	)

	user := &User{}
	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt); err != nil {
		return nil, err
	}

	return user, nil
}
