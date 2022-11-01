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

func (b *BadHabitPostgres) Create(userId int, input domain.BadHabitInput) (int, error) {
	tx, err := b.conn.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createItemQuery := fmt.Sprintf("insert into %s (habit_category_id, equivalent_id) values ($1, $2) returning id;", badHabitTable)
	row := tx.QueryRow(createItemQuery, input.HabitCategoryId, input.EquivalentId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, nil
	}
	createItemListQuery := fmt.Sprintf("insert into %s (user_id, bad_habit_id) values ($1, $2);", usersSpace)
	_, err = tx.Exec(createItemListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
func (b *BadHabitPostgres) GetAll(userId int) ([]domain.BadHabitOutput, error) {
	var items []domain.BadHabitOutput
	query := fmt.Sprintf("select bh.id, hc.name, eq.name, bh.created from %s as bh\n "+
		"inner JOIN %s as us on bh.id = us.bad_habit_id \n"+
		"inner JOIN %s as hc on bh.habit_category_id = hc.id\n"+
		"inner JOIN %s as eq on bh.equivalent_id = eq.id\n"+
		"where us.user_id = $1;", badHabitTable, usersSpace, categories, categories)
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
	query := fmt.Sprintf("select bh.id, hc.name, eq.name, bh.created from %s as bh\n "+
		"inner JOIN %s as us on bh.id = us.bad_habit_id \n"+
		"inner JOIN %s as hc on bh.habit_category_id = hc.id\n"+
		"inner JOIN %s as eq on bh.equivalent_id = eq.id\n"+
		"where us.user_id = $1 and bh.id = $2;", badHabitTable, usersSpace, categories, categories)
	err := b.conn.QueryRow(query, userId, id).Scan(&item.Id, &item.HabitCategory, &item.Equivalent, &item.Created)
	return item, err
}
func (b *BadHabitPostgres) DeleteHabitById(userId, id int) (string, error) {
	query := fmt.Sprintf("delete from %s bh using %s us where bh.id = us.bad_habit_id and us.user_id = $1 and us.bad_habit_id = $2", badHabitTable, usersSpace)
	_, err := b.conn.Exec(query, userId, id)
	if err != nil {
		return "", err
	}
	return "Record delete operation completed successfully", nil
}
func (b *BadHabitPostgres) GetCategory(categoryId, userId int) (string, error) {
	var valid Validate
	query := fmt.Sprintf("select bh.habit_category_id from %s as bd\n "+
		"inner JOIN %s as us on bh.id = us.bad_habit_id \n"+
		"inner JOIN %s as c on bh.habit_category_id = c.id\n"+
		"where us.user_id = $1;", badHabitTable, usersSpace, categories)
	err := b.conn.QueryRow(query, userId).Scan(&valid.HabitCategoryId)
	if valid.HabitCategoryId == categoryId {
		return negativeValidCategory, err
	}
	return positiveValidCategory, nil
}
func (b *BadHabitPostgres) GetCategories(offset, itemLimit int) ([]domain.HabitsCategory, error) {
	var categories []domain.HabitsCategory
	query := fmt.Sprintf("select c.id, c.name from categories as c where c.id between 1 and 8 limit $1 offset $2")
	row, err := b.conn.Query(query, itemLimit, offset)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var category domain.HabitsCategory
		err := row.Scan(&category.Id, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, err
}
func (b *BadHabitPostgres) GetEquivalents(offset, itemLimit int) ([]domain.Equivalents, error) {
	var equivalents []domain.Equivalents
	query := fmt.Sprintf("select c.id, c.name from categories as c where c.id between 9 and 18 limit $1 offset $2")
	row, err := b.conn.Query(query, itemLimit, offset)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var equivalent domain.Equivalents
		err := row.Scan(&equivalent.Id, &equivalent.Name)
		if err != nil {
			return nil, err
		}
		equivalents = append(equivalents, equivalent)
	}
	return equivalents, err
}
