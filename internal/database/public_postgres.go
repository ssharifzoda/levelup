package database

import (
	"fmt"
	"github.com/jackc/pgx"
	domain "github.com/ssharifzoda/levelup/internal/types"
	"gorm.io/gorm"
)

type PublicPostgres struct {
	conn    *pgx.Conn
	session *gorm.DB
}

func NewPublicPostgres(conn *pgx.Conn, session *gorm.DB) *PublicPostgres {
	return &PublicPostgres{conn: conn, session: session}
}
func (p *PublicPostgres) ReceivePublic(userId int, input domain.Public) error {
	tx := p.session.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	var id int
	publicQuery := fmt.Sprintf("insert into %s (age, gender_id, family_status_id, goal_to_life, big_fear)\n"+
		"values (?,?,?,?,?) returning id;", public)
	row := tx.Raw(publicQuery, input.Age, input.Gender, input.FamilyStatus, input.GoalToLife, input.BigFear)
	if err := row.Row().Scan(&id); err != nil {
		tx.Rollback()
		return err
	}
	query := fmt.Sprintf("update %s set public_id = ? where user_id = %d", usersSpace, userId)
	row = tx.Exec(query, id)
	if row.Error != nil {
		tx.Rollback()
		return row.Error
	}
	return tx.Commit().Error
}

func (p *PublicPostgres) DoTest(userId, temperamentId int) (string, error) {
	query := fmt.Sprintf("update %s set temperament_id = ? where user_id = %d", usersSpace, userId)
	row := p.session.Exec(query, temperamentId)
	if row.Error != nil {
		return row.Error.Error(), row.Error
	}
	var temperament string
	err := p.session.Table("categories").Select("name").Where("id = ?", temperamentId).Row().Scan(&temperament)
	if err != nil {
		return "", err
	}
	return temperament, nil
}
