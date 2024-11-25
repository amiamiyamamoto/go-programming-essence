package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type Todo struct {
	bun.BaseModel `bun:table:todos,alias:t`

	ID        int64     `bun:id,pk,autoincrement`
	Content   string    `bun:content,notnull`
	Done      bool      `bun:done`
	Until     time.Time `bun:until,nullzero`
	CreatedAt time.Time
	UpdatedAt time.Time `bun;,nilzero`
	DeletedAt time.Time `bun:",soft_delete,nullzero"`
}

func main() {
	sqldb, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer sqldb.Close()

	db := bun.NewDB(sqldb, pgdialect.New())
	defer db.Close()

	// TODO 抽出
	var todos []Todo
	ctx := context.Background()
	err = db.NewSelect().Model(&todos).Order("created_at").Where("until is not nill").Where("done is false").Scan(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
