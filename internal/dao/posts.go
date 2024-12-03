package dao

import (
	"myChat-API/internal/model"
)

func (d *DAO) GetThreadIdByUuid(uuid string) (int, error) {
	var thread_id int
	q := `SELECT id FROM threads WHERE uuid = $1;`
	row := d.DB.QueryRow(q, uuid)
	if err := row.Scan(&thread_id); err != nil {
		return 0, err
	}
	return thread_id, nil
}

func (d *DAO) SavePost(p model.Post) error {
	q := `INSERT INTO posts (uuid, body, thread_id, created_at) VALUES ($1, $2, $3, $4);`
	if _, err := d.DB.Exec(q, p.Uuid, p.Body, p.ThreadId, p.CreatedAt); err != nil {
		return err
	}
	return nil
}

func (d *DAO) GetPostSByThreadId(id int) ([]model.Post, error) {
	q := `SELECT uuid, body, created_at FROM posts WHERE thread_id = $1;`
	rows, err := d.DB.Query(q, id)
	if err != nil {
		return nil, err
	}

	res := make([]model.Post, 0)
	for rows.Next() {
		var p model.Post
		if err := rows.Scan(&p.Uuid, &p.Body, &p.CreatedAt); err != nil {
			return nil, err
		}
		res = append(res, p)
	}

	return res, nil
}
