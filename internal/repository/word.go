package repository

import "context"

type runeCount struct {
	character string `db:"character"`
	RuneCount int `db:"rune_count"`
}

func (r *Repository) GetRuneCount(ctx context.Context) ([]runeCount, error) {
	var rune_count []runeCount
	err := r.db.Select(&rune_count, "SELECT * FROM rone_counts")
	if err != nil {
		return nil, err
	}
	return rune_count, nil
}