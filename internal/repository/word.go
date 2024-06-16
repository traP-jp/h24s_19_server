package repository

import (
	"context"

	"github.com/gofrs/uuid"
)

type runeCount struct {
	RoomId    uuid.UUID `db:"room_id"`
	Character string    `db:"character"`
	RuneCount int       `db:"rune_count"`
}

func (r *Repository) GetRuneCount(roomId uuid.UUID) (map[rune]int, error) {
	var rune_count []runeCount
	err := r.db.Select(&rune_count, "SELECT * FROM `rune_counts` WHERE `room_id` = ?", roomId)
	if err != nil {
		return nil, err
	}
	runeCountMap := make(map[rune]int)
	for _, rc := range rune_count {
		runeCountMap[rune(rc.Character[0])] = rc.RuneCount
	}
	return runeCountMap, nil
}

func (r *Repository) UpdateRuneCount(ctx context.Context, roomId uuid.UUID, character string, RuneCount int) error {
	_, err := r.db.Exec("INSERT INTO `rune_counts` (`room_id`, `character`, `rune_count`) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE `rune_count` = ?", roomId, character, RuneCount, RuneCount)
	if err != nil {
		return err
	}
	return nil
}
