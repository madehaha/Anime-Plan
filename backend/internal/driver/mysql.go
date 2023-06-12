package driver

import (
	"backend/ent"
	"backend/internal/config"
	"backend/internal/logger"

	"context"
	"fmt"
	"go.uber.org/fx"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlClient(config config.AppConfig, lifecycle fx.Lifecycle) *ent.Client {
	host := config.Mysql.Host
	port := config.Mysql.Port
	user := config.Mysql.User
	password := config.Mysql.Password
	db := config.Mysql.DB // name
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", user, password, host, port, db)
	client, err := ent.Open("mysql", dataSourceName)
	if err != nil {
		info := fmt.Sprintf("Failed to open mysql, username: %s, password: %s, host: %s, port: %s, db: %s, %s, %v", user, password, host, port, db, dataSourceName, err)
		logger.Fatal(info)
	}

	err = client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf("Failed to create schema")
	}

	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return client.Close()
		}})
	return client
}
