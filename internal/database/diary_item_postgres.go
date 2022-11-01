package database

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/ssharifzoda/levelup/internal/types"
	"gorm.io/gorm"
)

type DiaryItemPostgres struct {
	conn    *pgx.Conn
	session *gorm.DB
}

func NewDiaryItemPostgres(conn *pgx.Conn, session *gorm.DB) *DiaryItemPostgres {
	return &DiaryItemPostgres{conn: conn, session: session}
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
	createItemListQuery := fmt.Sprintf("insert into %s (user_id, item_id) values ($1, $2)", usersSpace)
	_, err = tx.Exec(createItemListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
func (d *DiaryItemPostgres) GetAll(userId, offset, itemLimit int) ([]domain.Item, error) {
	var items []domain.Item
	query := fmt.Sprintf("select it.id, it.title, it.description, it.body from %s it\n"+
		"inner join %s us on it.id = us.item_id where us.user_id = $1 limit $2 offset $3", itemTable, usersSpace)
	row, err := d.conn.Query(query, userId, itemLimit, offset)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var item domain.Item
		err = row.Scan(&item.Id, &item.Title, &item.Description, &item.Body)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
func (d *DiaryItemPostgres) GetById(userId, itemId int) (domain.Item, error) {
	var item domain.Item
	var list domain.UsersSpace
	query := fmt.Sprintf("select it.id, it.title, it.description, it.body, us.id, us.user_id, us.item_id from %s it inner join %s us on it.id = us.item_id where us.user_id = $1 and item_id = $2", itemTable, usersSpace)
	err := d.conn.QueryRow(query, userId, itemId).Scan(&item.Id, &item.Title, &item.Description, &item.Body, &list.UserId, &list.DiaryId)
	return item, err
}
func (d *DiaryItemPostgres) DeleteItemById(userId, itemId int) (string, error) {
	query := fmt.Sprintf("delete from %s tl using %s ul where tl.id = ul.item_id and ul.user_id = $1 and ul.item_id = $2", itemTable, usersSpace)
	_, err := d.conn.Exec(query, userId, itemId)
	if err != nil {
		return "", err
	}
	return "Record delete operation completed successfully", nil
}
func (d *DiaryItemPostgres) GetItemByTitle(userId int, title string) (domain.Item, error) {
	var item domain.Item
	err := d.session.Table("item").Select("item.id,title, description, body, created").Joins("join users_space us on user_id = ?", userId).Where("item.title = ?", title).Find(&item)
	if err.Error != nil {
		return item, err.Error
	}
	return item, nil
}
