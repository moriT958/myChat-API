package dao

import "myChat-API/internal/model"

func (d *DAO) SaveUser(u model.User) error {
	q := `INSERT INTO users (uuid, username, password, created_at) VALUES ($1, $2, $3, $4);`

	_, err := d.DB.Exec(q, u.Uuid, u.Username, u.Password, u.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (d *DAO) GetUserByUsername(username string) (model.User, error) {
	var u model.User
	q := `SELECT FROM users WHERE username = $1;`
	row := d.DB.QueryRow(q, username)
	if err := row.Scan(&u.Id, &u.Uuid, &u.Username, &u.Password, &u.CreatedAt); err != nil {
		return model.User{}, err
	}
	return u, nil
}
