package model

import (
	"context"
	"time"

	bunsample "github.com/nakaaaa/bun-sample"
	"github.com/uptrace/bun"
)

type Story struct {
	bun.BaseModel `bun:"table:stories"`
	ID            uint64    `bun:",pk"`
	Title         string    `bun:",notnull"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     bun.NullTime
}

func (s *Story) TableName() string {
	return "stories"
}

func CreateTableStory(ctx context.Context, db *bunsample.DB, story Story) error {
	_, err := db.NewCreateTable().
		Model(&story).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
