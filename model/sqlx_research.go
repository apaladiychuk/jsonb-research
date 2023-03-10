package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type SqlxRepository struct {
	conn *sqlx.DB
}

func NewSqlXRepo() (*SqlxRepository, error) {
	conn, err := sqlx.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &SqlxRepository{
		conn: conn,
	}, nil
}
func (r *SqlxRepository) GetAll() ([]*Security, error) {
	start := time.Now()
	var result []*Security
	if err := r.conn.Select(&result, `select * from security`); err != nil {
		return nil, err
	}
	fmt.Printf(" sqlx result %d\n", time.Now().Sub(start))
	return result, nil
}

func (r *SqlxRepository) GetAllColumn() ([]*Security, error) {
	start := time.Now()
	var result []*Security
	if err := r.conn.Select(&result, `select id,name  from security`); err != nil {
		return nil, err
	}
	fmt.Printf(" sqlx  GetAllColumn result %d\n", time.Now().Sub(start))
	return result, nil
}
