package config

import (
	"context"
	"database/sql"
	"fmt"

	go_ora "github.com/sijms/go-ora/v2"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/oracledialect"
	"github.com/uptrace/bun/extra/bundebug"

	"bun-spreader/models"
	"bun-spreader/utils"
)

var DB *bun.DB

func Init() error {
	port := 1521
	connStr := go_ora.BuildUrl("localhost", port, "freepdb1", "narie", "password", nil)
	conn, err := sql.Open("oracle", connStr)
	utils.HandleError(err)

	DB = bun.NewDB(conn, oracledialect.New())

	DB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))
	ctx := context.Background()

	res, err := DB.NewCreateTable().Model((*models.Customer)(nil)).Exec(ctx)
	utils.HandleError(err)
	fmt.Println(res)

	/* 	err = DB.ResetModel(ctx, &models.Customer{}) */

	return nil
}
