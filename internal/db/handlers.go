package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/uptrace/bun"
)

func GetMovies(ctx context.Context) ([]Movie, error) {
	var m []Movie
	if err := client.NewSelect().Model(&m).Scan(ctx); err != nil {
		return nil, err
	}
	return m, nil
}

func GetMovie(ctx context.Context, m Movie) (*Movie, error) {
	q := client.NewSelect().Model(&m)

	if m.ID != 0 {
		q.WhereOr("id ? ?", bun.Safe("="), m.ID)
	}
	if m.Name != "" {
		q.WhereOr("? ILIKE ?", bun.Ident("name"), fmt.Sprintf("%%%s%%", m.Name))
	}
	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return &m, nil
}

func UpsertMovie(ctx context.Context, m *Movie) error {
	if m.Name == "" {
		return errors.New("movie name is required but not provided")
	}
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
