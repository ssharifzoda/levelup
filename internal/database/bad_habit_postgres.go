package database

import (
	"fmt"
	"github.com/jackc/pgx"
	domain "github.com/ssharifzoda/levelup/internal/types"
)

type BadHabitPostgres struct {
	conn *pgx.Conn
}

func NewBadHabitPostgres(conn *pgx.Conn) *BadHabitPostgres {
	return &BadHabitPostgres{conn: conn}
}

func (b *BadHabitPostgres) Create(userId int, input domain.BadHabit) (int, error) {
	tx, err := b.conn.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createItemQuery := fmt.Sprintf("insert into %s (bad_habit, equivalent) values ($1, $2) returning id", badHabitTable)
	row := tx.QueryRow(createItemQuery, input.BadHabit, input.Equivalent)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, nil
	}
	createItemListQuery := fmt.Sprintf("insert into %s (user_id, bad_habit_id) values ($1, $2)", badHabitList)
	_, err = tx.Exec(createItemListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
