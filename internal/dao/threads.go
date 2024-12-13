package dao

import "myChat-API/internal/model"

func (d *DAO) SaveThread(t model.Thread) error {
	q := `INSERT INTO threads (uuid, topic, created_at, user_id) VALUES ($1, $2, $3, $4);`

	_, err := d.DB.Exec(q, t.Uuid, t.Topic, t.CreatedAt, t.UserId)
	if err != nil {
		return err
	}

	return nil
}

func (d *DAO) GetThreads(limit int, offset int) ([]model.Thread, error) {

	sql := `SELECT uuid, topic, created_at FROM threads LIMIT $1 OFFSET $2;`
	rows, err := d.DB.Query(sql, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]model.Thread, 0)
	for rows.Next() {
		var t model.Thread
		if err := rows.Scan(&t.Uuid, &t.Topic, &t.CreatedAt); err != nil {
			return nil, err
		}
		res = append(res, t)
	}

	return res, nil
}

func (d *DAO) GetThreadByUuid(uuid string) (model.Thread, error) {
	var t model.Thread
	q := `SELECT id, uuid, topic, created_at FROM threads WHERE uuid = $1;`
	row := d.DB.QueryRow(q, uuid)
	if err := row.Scan(&t.Id, &t.Uuid, &t.Topic, &t.CreatedAt); err != nil {
		return model.Thread{}, err
	}
	return t, nil
}
