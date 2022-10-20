package database

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/ssharifzoda/levelup/internal/types"
)

type DiaryItemPostgres struct {
	conn *pgx.Conn
}

func NewDiaryItemPostgres(conn *pgx.Conn) *DiaryItemPostgres {
	return &DiaryItemPostgres{conn: conn}
}
func (d *DiaryItemPostgres) Create(userId int, item domain.Item) (int, error) {
	tx, err := d.conn.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createItemQuery := fmt.Sprintf("insert into %s (title, description, body) values ($1, $2, $3) returning id", itemTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description, item.Body)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, nil
	}
	createItemListQuery := fmt.Sprintf("insert into %s (user_id, item_id) values ($1, $2)", itemsListTable)
	_, err = tx.Exec(createItemListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
func (d *DiaryItemPostgres) GetAll(userId int) ([]domain.Item, error) {
	var items []domain.Item
	query := fmt.Sprintf("select tl.id, tl.title, tl.description, tl.body, ul.id, ul.user_id, ul.item_id from %s tl inner join %s ul on tl.id = ul.item_id where ul.user_id = $1", itemTable, itemsListTable)
	row, err := d.conn.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var item domain.Item
		var list domain.ItemList
		err = row.Scan(&item.Id, &item.Title, &item.Description, &item.Body, &list.Id, &list.UserId, &list.DiaryId)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
func (d *DiaryItemPostgres) GetById(userId, itemId int) (domain.Item, error) {
	var item domain.Item
	var list domain.ItemList
	query := fmt.Sprintf("select tl.id, tl.title, tl.description, tl.body, ul.id, ul.user_id, ul.item_id from %s tl inner join %s ul on tl.id = ul.item_id where ul.user_id = $1 and item_id = $2", itemTable, itemsListTable)
	err := d.conn.QueryRow(query, userId, itemId).Scan(&item.Id, &item.Title, &item.Description, &item.Body, &list.Id, &list.UserId, &list.DiaryId)
	return item, err
}
func (d *DiaryItemPostgres) DeleteItemById(userId, itemId int) (string, error) {
	query := fmt.Sprintf("delete from %s tl using %s ul where tl.id = ul.item_id and ul.user_id = $1 and ul.item_id = $2", itemTable, itemsListTable)
	_, err := d.conn.Exec(query, userId, itemId)
	if err != nil {
		return "", err
	}
	return "Record delete operation completed successfully", nil
}
