package database

import (
	"fmt"
	"github.com/jackc/pgx"
	domain "github.com/ssharifzoda/levelup/internal/types"
)

type MentalDevelopPostgres struct {
	conn *pgx.Conn
}

func NewMentalDevelopPostgres(conn *pgx.Conn) *MentalDevelopPostgres {
	return &MentalDevelopPostgres{conn: conn}
}

func (m *MentalDevelopPostgres) Create(userId int, input domain.CourseInput) (int, error) {
	tx, err := m.conn.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createItemQuery := fmt.Sprintf("insert into %s (mental_category_id) values ($1) returning id", mentalCourseTable)
	row := tx.QueryRow(createItemQuery, input.MentalCategoryId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, nil
	}
	createItemListQuery := fmt.Sprintf("insert into %s (user_id, mental_course_id) values ($1, $2)", usersSpace)
	_, err = tx.Exec(createItemListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
func (m *MentalDevelopPostgres) GetById(userId int, id int) (domain.CourseOutput, error) {
	var item domain.CourseOutput
	query := fmt.Sprintf("select mc.id, c.name, mc.created from %s as mc\n"+
		"inner JOIN %s as us on mc.id = us.mental_course_id \n"+
		"inner join %s as c on mc.mental_category_id = c.id\n"+
		"where us.user_id = $1 and mc.id = $2", mentalCourseTable, usersSpace, categories)
	err := m.conn.QueryRow(query, userId, id).Scan(&item.Id, &item.MentalCategory, &item.Created)
	return item, err
}
func (m *MentalDevelopPostgres) DeleteCourseById(userId, id int) (string, error) {
	query := fmt.Sprintf("delete from %s c using %s mc where c.id = mc.course_id and mc.user_id = $1 and mc.course_id = $2", mentalCourseTable, usersSpace)
	_, err := m.conn.Exec(query, userId, id)
	if err != nil {
		return "", err
	}
	return "Record delete operation completed successfully", nil
}
func (m *MentalDevelopPostgres) GetCategory(categoryId, userId int) (string, error) {
	var valid Validate
	query := fmt.Sprintf("select c.mental_category_id from %s as c\n "+
		"inner JOIN %s as mcl on c.id = mcl.course_id \n"+
		"inner JOIN %s as mc on c.mental_category_id = mc.id\n"+
		"where mcl.user_id = $1;", mentalCourseTable, usersSpace, categories)
	err := m.conn.QueryRow(query, userId).Scan(&valid.CourseCategory)
	if valid.CourseCategory == categoryId {
		return negativeValidCategory, err
	}
	return positiveValidCategory, nil
}
func (m *MentalDevelopPostgres) GetCategories() ([]domain.MentalCourseCategory, error) {
	var categories []domain.MentalCourseCategory
	query := fmt.Sprintf("select c.id, c.name from categories as c where c.id between 19 and 22")
	row, err := m.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var category domain.MentalCourseCategory
		err := row.Scan(&category.Id, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, err
}
