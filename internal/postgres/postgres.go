package postgres

import (
	"context"
	"github.com/jackc/pgx"
	"time"
)

const (
	usersTable     = "users"
	diaryListTable = "diary_list"
	userListTable  = "user_list"
	diaryItemTable = "diary_items"
	itemsListTable = "items_list"
)

type Config struct {
	Host     string
	Port     uint16
	Username string
	Password string
	DBName   string
}

func NewPostgresDB(cfg Config) (*pgx.Conn, error) {
	conf := pgx.ConnConfig{Host: cfg.Host, Port: cfg.Port, User: cfg.Username, Password: cfg.Password, Database: cfg.DBName}
	conn, err := pgx.Connect(conf)
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
