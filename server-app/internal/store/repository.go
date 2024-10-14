package store

import (
	"fmt"
	"mtg/internal/model"
)

type Rep struct {
	storage *Storage
}

func (r *Rep) CreateData(s *model.Data) error {
	q := `INSERT INTO data_table (name, age, isconnected) VALUES ($1, $2, $3) RETURNING id`

	r.storage.logger.Debugf(fmt.Sprintf("SQL Query: %s", q))

	if err := r.storage.db.QueryRow(
		q,
		s.Name,
		s.Age,
		s.IsConnected,
	).Scan(&s.ID); err != nil {
		r.storage.logger.Errorf("Ошибка SQL: %s", err)
		return err
	}
	return nil
}

func (r *Rep) GetData() ([]*model.Data, error) {
	q := `SELECT id, name, age, isconnected
	 FROM data_table`

	r.storage.logger.Debugf(fmt.Sprintf("SQL Query: %s", q))

	rows, err := r.storage.db.Query(q)
	if err != nil {
		r.storage.logger.Errorf("Ошибка SQL: %s", err)
		return nil, err
	}

	var dataArr []*model.Data
	for rows.Next() {
		d := new(model.Data)
		err := rows.Scan(
			&d.ID,
			&d.Name,
			&d.Age,
			&d.IsConnected,
		)
		if err != nil {
			return nil, err
		}
		dataArr = append(dataArr, d)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return dataArr, nil
}
