package database

import (
	"fmt"
	"github.com/jackc/pgx"
	domain "github.com/ssharifzoda/levelup/internal/types"
)

type PhysicianDevelopPostgres struct {
	conn *pgx.Conn
}

func NewPhysicianDevelopPostgres(conn *pgx.Conn) *PhysicianDevelopPostgres {
	return &PhysicianDevelopPostgres{conn: conn}
}

func (p *PhysicianDevelopPostgres) Create(userId int, input domain.BodyCourseInput) (int, error) {
	tx, err := p.conn.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createItemQuery := fmt.Sprintf("insert into %s (level_id, train_category_id) values ($1, $2) returning id", bodyCourseTable)
	row := tx.QueryRow(createItemQuery, input.LevelId, input.TrainCategoryId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, nil
	}
	createItemListQuery := fmt.Sprintf("insert into %s (user_id, body_course_id) values ($1, $2)", bodyCourselist)
	_, err = tx.Exec(createItemListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
func (p *PhysicianDevelopPostgres) GetById(userId int, id int) (domain.BodyCourseOutput, error) {
	var item domain.BodyCourseOutput
	query := fmt.Sprintf("select bc.id, lc.name, tc.name, bc.created from %s as bc\n"+
		"inner JOIN %s as bcl on bc.id = bcl.body_course_id \n"+
		"inner join %s as lc on bc.level_id = lc.id\n"+
		"inner join %s as tc on bc.train_category_id = tc.id\n"+
		"where bcl.user_id = $1 and bc.id = $2;", bodyCourseTable, bodyCourselist, bodyLevelTable, trainCategoryTable)
	err := p.conn.QueryRow(query, userId, id).Scan(&item.Id, &item.Level, &item.TrainCategory, &item.Created)
	return item, err
}
