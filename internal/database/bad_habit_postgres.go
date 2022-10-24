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

type ValidateCategory struct {
	HabitCategoryId int `json:"habit_category_id"`
}

func (b *BadHabitPostgres) Create(userId int, input domain.BadHabitInput) (int, error) {
	tx, err := b.conn.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createItemQuery := fmt.Sprintf("insert into %s (habit_category_id, equivalent_id) values ($1, $2) returning id", badHabitTable)
	row := tx.QueryRow(createItemQuery, input.HabitCategoryId, input.EquivalentId)
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
func (b *BadHabitPostgres) GetAll(userId int) ([]domain.BadHabitOutput, error) {
	var items []domain.BadHabitOutput
	query := fmt.Sprintf("select bd.id, hc.name, eq.name, bd.created from %s as bd\n "+
		"inner JOIN %s as ul on bd.id = ul.bad_habit_id \n"+
		"inner JOIN %s as hc on bd.habit_category_id = hc.id\n"+
		"inner JOIN %s as eq on bd.equivalent_id = eq.id\n"+
		"where ul.user_id = $1;", badHabitTable, badHabitList, badHabitCategory, equivalents)
	row, err := b.conn.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var item domain.BadHabitOutput
		err = row.Scan(&item.Id, &item.HabitCategory, &item.Equivalent, &item.Created)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
func (b *BadHabitPostgres) GetById(userId, id int) (domain.BadHabitOutput, error) {
	var item domain.BadHabitOutput
	query := fmt.Sprintf("select bd.id, hc.name, eq.name, bd.created from %s as bd\n "+
		"inner JOIN %s as ul on bd.id = ul.bad_habit_id \n"+
		"inner JOIN %s as hc on bd.habit_category_id = hc.id\n"+
		"inner JOIN %s as eq on bd.equivalent_id = eq.id\n"+
		"where ul.user_id = $1 and bd.id = $2;", badHabitTable, badHabitList, badHabitCategory, equivalents)
	err := b.conn.QueryRow(query, userId, id).Scan(&item.Id, &item.HabitCategory, &item.Equivalent, &item.Created)
	return item, err
}
func (b *BadHabitPostgres) DeleteHabitById(userId, id int) (string, error) {
	query := fmt.Sprintf("delete from %s bd using %s ul where bd.id = ul.bad_habit_id and ul.user_id = $1 and ul.bad_habit_id = $2", badHabitTable, badHabitList)
	_, err := b.conn.Exec(query, userId, id)
	if err != nil {
		return "", err
	}
	return "Record delete operation completed successfully", nil
}
func (b *BadHabitPostgres) GetCategory(categoryId, userId int) (string, error) {
	var valid ValidateCategory
	query := fmt.Sprintf("select bd.habit_category_id from %s as bd\n "+
		"inner JOIN %s as ul on bd.id = ul.bad_habit_id \n"+
		"inner JOIN %s as hc on bd.habit_category_id = hc.id\n"+
		"where ul.user_id = $1;", badHabitTable, badHabitList, badHabitCategory)
	err := b.conn.QueryRow(query, userId).Scan(&valid.HabitCategoryId)
	if valid.HabitCategoryId == categoryId {
		return negativeValidCategory, err
	}
	return positiveValidCategory, nil
}
