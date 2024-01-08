package db

import (
	"context"
	"time"
)

func GetMovie(ctx context.Context, id int64) (*Movie, error) {
	var m Movie
	if err := client.NewSelect().Model(&m).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return &m, nil
}

func CreateMovie(ctx context.Context, m *Movie) error {
	m.AddedAt = time.Now()
	_, err := client.NewInsert().
		Model(m).
		On("CONFLICT (id) DO UPDATE").
		Set("added_at = EXCLUDED.added_at").
		Set("digitized_at = EXCLUDED.digitized_at").
		Exec(ctx)

	if err != nil {
		return err
	}
	return nil
}
