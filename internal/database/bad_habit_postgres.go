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
	input.Session = 1
	createItemQuery := fmt.Sprintf("insert into %s (bad_habit, equivalent, session) values ($1, $2, $3) returning id", badHabitTable)
	row := tx.QueryRow(createItemQuery, input.BadHabit, input.Equivalent, input.Session)
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
func (b *BadHabitPostgres) GetAll(userId int) ([]domain.BadHabit, error) {
	var items []domain.BadHabit
	query := fmt.Sprintf("select tl.id, tl.bad_habit, tl.equivalent, tl.session from %s tl inner join %s ul on tl.id = ul.bad_habit_id where ul.user_id = $1", badHabitTable, badHabitList)
	row, err := b.conn.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var item domain.BadHabit
		err = row.Scan(&item.Id, &item.BadHabit, &item.Equivalent, &item.Session)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
func (b *BadHabitPostgres) GetById(userId, id int) (domain.BadHabit, error) {
	var item domain.BadHabit
	query := fmt.Sprintf("select tl.id, tl.bad_habit, tl.equivalent, tl.session from %s tl inner join %s ul on tl.id = ul.bad_habit_id where ul.user_id = $1 and bad_habit_id = $2", badHabitTable, badHabitList)
	err := b.conn.QueryRow(query, userId, id).Scan(&item.Id, &item.BadHabit, &item.Equivalent, &item.Session)
	return item, err
}
func (b *BadHabitPostgres) DeleteHabitById(userId, id int) (string, error) {
	query := fmt.Sprintf("delete from %s tl using %s ul where tl.id = ul.bad_habit_id and ul.user_id = $1 and ul.bad_habit_id = $2", badHabitTable, badHabitList)
	_, err := b.conn.Exec(query, userId, id)
	if err != nil {
		return "", err
	}
	return "Record delete operation completed successfully", nil
}
func (b *BadHabitPostgres) DoExercise(userId, id int, input domain.DoExercise) (string, error) {
	query := fmt.Sprintf("update %s tl set session =session+%d, last_session = $3 from %s ul where tl.id = ul.bad_habit_id and ul.user_id = $1 and ul.bad_habit_id = $2", badHabitTable, input.Session, badHabitList)
	_, err := b.conn.Exec(query, userId, id, input.LastSession)
	if err != nil {
		return "", err
	}
	return "Record update operation completed successfully", nil
}
