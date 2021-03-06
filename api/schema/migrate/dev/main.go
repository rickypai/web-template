package main

import (
	"context"
	"log"

	"github.com/rickypai/web-template/api/config"
	"github.com/rickypai/web-template/api/schema/migrate/common"
)

func main() {
	ctx := context.Background()

	db, err := config.DB(ctx)
	if err != nil {
		log.Fatalf("creating database connection: %s", err)
	}

	err = common.MigrateInstance(ctx, db)
	if err != nil {
		log.Fatalf("running migrations: %s", err)
	}
}
