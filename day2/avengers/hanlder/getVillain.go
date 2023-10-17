package hanlder

import (
	"avengers2/entity"
	"context"
	"database/sql"
	"time"
)

func GetVillain(db *sql.DB) ([]entity.Villain, error) {
	// Heroes Object
	villain := []entity.Villain{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT name, universe, imageURL FROM villain
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v entity.Villain

		err := rows.Scan(&v.Name, &v.Universe, &v.ImageURL)
		if err != nil {
			return nil, err
		}

		villain = append(villain, v)
	}

	return villain, nil
}
