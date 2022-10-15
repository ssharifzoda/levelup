package database

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/ssharifzoda/levelup/internal/domain"
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
