package main

import (
	"context"

	bunsample "github.com/nakaaaa/bun-sample"
	"github.com/nakaaaa/bun-sample/model"
)

func main() {
	ctx := context.Background()
	db := bunsample.DB{}
	if err := db.Open(); err != nil {
		panic(err)
	}

	story := model.Story{}

	if err := model.CreateTableStory(ctx, &db, story); err != nil {
		panic(err)
	}
}
