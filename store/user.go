package store

import "time"

func (s *Store) CreateUser(userName string, creationTime time.Time) (int64, error) {
	stmt, err := s.db.Prepare(`
	  INSERT INTO Users (UserName, CreationTime)
	  VALUES (?, ?)
	`)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(userName, creationTime.Format(time.RFC3339))
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}
