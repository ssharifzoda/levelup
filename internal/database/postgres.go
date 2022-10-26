package database

import (
	"context"
	"github.com/jackc/pgx"
	"time"
)

const (
	usersTable            = "users"
	itemTable             = "item"
	itemsListTable        = "items_list"
	badHabitTable         = "bad_habit"
	badHabitList          = "bad_habits_list"
	badHabitCategory      = "habits_category"
	equivalents           = "equivalents"
	positiveValidCategory = "He did not have this bad habits"
	negativeValidCategory = "You already have this bad habits"
	courseTable           = "course"
	mentalCategoryTable   = "mental_category"
	mentalCourseList      = "mental_course_list"
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