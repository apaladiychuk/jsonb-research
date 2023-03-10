package model

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"time"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	if query, err := q.FormattedQuery(); err != nil {
		fmt.Println("[SQL]: " + err.Error())
	} else {
		fmt.Println(string(query))
	}

	return nil
}

func NewRepo() *SecurityRepository {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
	})
	db.AddQueryHook(dbLogger{})
	return &SecurityRepository{
		db: db,
	}
}

func (r *SecurityRepository) GetAll() ([]*Security, error) {
	start := time.Now()
	var result []*Security
	if err := r.db.Model(&Security{}).Select(&result); err != nil {
		return nil, err
	}
	fmt.Printf(" go-pg/pg  result %d\n ", time.Now().Sub(start))
	return result, nil
}

func (r *SecurityRepository) Create(s *Security) error {
	if _, err := r.db.Model(s).Insert(s); err != nil {
		return err
	}
	return nil
}
func (r *SecurityRepository) Update(s *Security) error {
	if _, err := r.db.Model(s).Update(s); err != nil {
		return err
	}
	return nil
}
func (r *SecurityRepository) GetAllColumns() ([]*Security, error) {
	start := time.Now()
	var result []*Security
	if err := r.db.Model(&Security{}).
		Column("id", "name").
		Select(&result); err != nil {
		return nil, err
	}
	fmt.Printf(" go-pg/pg   GetAllColumns result %d\n ", time.Now().Sub(start))
	return result, nil
}
