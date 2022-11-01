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
	createItemListQuery := fmt.Sprintf("insert into %s (user_id, body_course_id) values ($1, $2)", usersSpace)
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
		"where bcl.user_id = $1 and bc.id = $2;", bodyCourseTable, usersSpace, categories, categories)
	err := p.conn.QueryRow(query, userId, id).Scan(&item.Id, &item.Level, &item.TrainCategory, &item.Created)
	return item, err
}

func (p *PhysicianDevelopPostgres) DeleteCourseById(userId int, id int) (string, error) {
	query := fmt.Sprintf("delete from %s bc using %s bcl where bc.id = bcl.body_course_id and bcl.user_id = $1 and bcl.body_course_id = $2", bodyCourseTable, usersSpace)
	_, err := p.conn.Exec(query, userId, id)
	if err != nil {
		return "", err
	}
	return "Record delete operation completed successfully", nil
}

func (p *PhysicianDevelopPostgres) GetCategory(trainCategoryId, userId int) (string, error) {
	var valid Validate
	query := fmt.Sprintf("select bc.train_category_id from %s as bc\n "+
		"inner JOIN %s as bcl on bc.id = bcl.body_course_id \n"+
		"inner JOIN %s as tc on bc.train_category_id = tc.id\n"+
		"where bcl.user_id = $1;", bodyCourseTable, usersSpace, categories)
	err := p.conn.QueryRow(query, userId).Scan(&valid.TrainCategoryId)
	if valid.TrainCategoryId == trainCategoryId {
		return negativeValidCategory, err
	}
	return positiveValidCategory, nil
}
func (p *PhysicianDevelopPostgres) GetCategories() ([]domain.BodyCourseCategories, error) {
	var categories []domain.BodyCourseCategories
	query := fmt.Sprintf("select c.id, c.name from categories as c where c.id between 26 and 30")
	row, err := p.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var category domain.BodyCourseCategories
		err := row.Scan(&category.Id, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, err
}
func (p *PhysicianDevelopPostgres) GetLevels() ([]domain.BodyLevelCourse, error) {
	var levels []domain.BodyLevelCourse
	query := fmt.Sprintf("select c.id, c.name from categories as c where c.id between 23 and 25")
	row, err := p.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var level domain.BodyLevelCourse
		err := row.Scan(&level.Id, &level.Name)
		if err != nil {
			return nil, err
		}
		levels = append(levels, level)
	}
	return levels, err
}
