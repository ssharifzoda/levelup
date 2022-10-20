package database

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/ssharifzoda/levelup/internal/types"
)

type AuthPostgres struct {
	conn *pgx.Conn
}

func NewAuthPostgres(conn *pgx.Conn) *AuthPostgres {
	return &AuthPostgres{conn: conn}
}
func (a *AuthPostgres) CreateUser(user domain.User) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (name, username, password_hash) values ($1, $2, $3) returning id", usersTable)
	row := a.conn.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (a *AuthPostgres) GetUser(username, password string) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("select id, username, password_hash from %s where username=$1 and password_hash=$2", usersTable)
	row := a.conn.QueryRow(query, username, password)
	err := row.Scan(&user.Id, &user.Username, &user.Password)
	return user, err
}
