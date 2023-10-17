package hanlder

import (
	"avengers2/entity"
	"context"
	"database/sql"
	"time"
)

func GetHeroes(db *sql.DB) ([]entity.Heroes, error) {
	// Heroes Object
	heroes := []entity.Heroes{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT name, universe, skill, imageURL FROM heroes
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var h entity.Heroes

		err := rows.Scan(&h.Name, &h.Universe, &h.Skill, &h.ImageURL)
		if err != nil {
			return nil, err
		}

		heroes = append(heroes, h)
	}

	return heroes, nil
}
