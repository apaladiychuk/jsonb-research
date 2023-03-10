package model

import (
	pg "github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type Security struct {
	tableName struct{}  `pg:"security,alias:s"`
	ID        uuid.UUID `json:"id" pg:"id,pk"`
	Name      string    `json:"name" pg:"name"`
	Metadata  *Metadata `json:"metadata" pg:"metadata"`
}

type SecurityRepository struct {
	db *pg.DB
}

type Sukuk struct {
	WSID   string `json:"wsid"`
	Param1 string `json:"param1"`
}

type ExternalSecurity struct {
	ISIN   string `json:"isin"`
	Param2 string `json:"param2"`
}
