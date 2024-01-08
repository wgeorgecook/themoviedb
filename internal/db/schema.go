package db

import (
	"time"

	"github.com/uptrace/bun"
)

type Movie struct {
	bun.BaseModel `bun:"table:movies,alias:m"`

	ID          int64     `bun:",pk,autoincrement"`
	Name        string    `bun:"name"`
	AddedAt     time.Time `bun:"added_at"`
	DigitizedAt time.Time `bun:"digitized_at"`
}
