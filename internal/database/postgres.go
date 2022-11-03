package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/spf13/viper"
	"github.com/ssharifzoda/levelup/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

const (
	usersTable            = "users"
	itemTable             = "item"
	badHabitTable         = "bad_habit"
	positiveValidCategory = "He did not have this"
	negativeValidCategory = "You already have this"
	mentalCourseTable     = "mental_course"
	bodyCourseTable       = "body_course"
	usersSpace            = "users_space"
	categories            = "categories"
	public                = "public"
)

type Validate struct {
	HabitCategoryId int `json:"habit_category_id"`
	CourseCategory  int `json:"mental_category_id"`
	TrainCategoryId int `json:"train_category_id"`
}

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
func NewPostgresGorm() (*gorm.DB, error) {
	Host := viper.GetString("db.host")
	Port := viper.GetUint16("db.port")
	Username := viper.GetString("db.username")
	Password := os.Getenv("DB_PASSWORD")
	DBName := viper.GetString("db.dbname")
	logger := logging.GetLogger()
	connString := fmt.Sprintf("host=%s user=%s password=%v dbname=%s port=%d sslmode=disable TimeZone=Asia/Dushanbe",
		Host, Username, Password, DBName, Port)
	conn, err := gorm.Open(postgres.Open(connString))
	if err != nil {
		logger.Printf("%s GetPostgresConnection -> Open error: ", err.Error())
		return nil, err
	}

	logger.Println("Postgres Connection success: ", Host)

	return conn, nil
}
