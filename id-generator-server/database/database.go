package database

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"
)

var db *sql.DB

//const dbName = "uniquekeys.db"

type UniqueKey struct {
	Key      string
	IsBooked bool
	IsUsed   bool
}

func addUniqueKey(uk *UniqueKey) error {
	_, err := db.ExecContext(
		context.Background(),
		`INSERT INTO unique_key (key, isBooked, isUsed) VALUES (?,?,?);`, uk.Key, uk.IsBooked, uk.IsUsed,
	)
	// rest of the function
	if err != nil {
		return err
	}

	return nil
}

func setIsBooked(key string, isBooked bool) error {
	_, err := db.ExecContext(
		context.Background(),
		`UPDATE unique_key SET isBooked=? WHERE key=?;`, isBooked, key,
	)
	// rest of the function
	if err != nil {
		return err
	}

	return nil
}

func setIsUsed(key string, isUsed bool) error {
	_, err := db.ExecContext(
		context.Background(),
		`UPDATE unique_key SET isUsed=? WHERE key=?;`, isUsed, key,
	)

	if err != nil {
		return err
	}

	return nil
}

func getKey(key string) (UniqueKey, error) {

	var uniqueKey UniqueKey

	row := db.QueryRowContext(
		context.Background(),
		`SELECT * FROM unique_key WHERE key=?`, key,
	)

	err := row.Scan(&uniqueKey.Key, &uniqueKey.IsBooked, &uniqueKey.IsUsed)

	if err != nil {
		return UniqueKey{}, err
	}

	return uniqueKey, nil
}

func getKeys(count int, isBooked bool, isUsed bool) ([]UniqueKey, error) {
	var keys []UniqueKey
	rows, err := db.QueryContext(
		context.Background(),
		`SELECT * FROM unique_key WHERE isBooked=? AND isUsed=? LIMIT ?;`, isBooked, isUsed, count,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var uniqueKey UniqueKey

		if err := rows.Scan(
			&uniqueKey.Key, &uniqueKey.IsBooked, &uniqueKey.IsUsed,
		); err != nil {
			return nil, err
		}
		keys = append(keys, uniqueKey)
	}
	return keys, err
}

func initDatabase(dbPath string) error {
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}

	_, err = db.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS unique_key (
			key TEXT PRIMARY KEY,
			isBooked NUMERIC NOT NULL,
			isUsed NUMERIC NOT NULL
		)`,
	)

	if err != nil {
		return err
	}
	return nil
}
