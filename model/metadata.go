package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const DocTypeSukuk = "sukuk"
const DocTypeExternalSecurity = "external_security"

type Metadata struct {
	DocumentType string
	Version      string
	Data         JSONMap
}

func (m Metadata) Value() (driver.Value, error) {
	j, err := json.Marshal(m)
	return string(j), err
}

func (m *Metadata) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("type assertion .([]byte) failed")
	}
	err := json.Unmarshal(source, m)
	if err != nil {
		return err
	}
	if m == nil {
		err = json.Unmarshal([]byte("{}"), m)
		if err != nil {
			return err
		}
	}
	return nil
}

type JSONMap map[string]interface{}

func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return json.Marshal(make(JSONMap))
	}
	return json.Marshal(j)
}

func (j *JSONMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("type assertion .([]byte) failed")
	}
	err := json.Unmarshal(source, j)
	if err != nil {
		return err
	}
	if *j == nil {
		err = json.Unmarshal([]byte("{}"), j)
		if err != nil {
			return err
		}
	}
	return nil
}

func (j *JSONMap) ToStruct(dest interface{}) error {
	buf, err := json.Marshal(j)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, &dest)
}

func (j *JSONMap) FromStruct(src interface{}) error {
	buf, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, j)
}

type Wrapper[T any] struct {
	docType string
}

func NewWrapper[T any](docType string) *Wrapper[T] {
	return &Wrapper[T]{docType: docType}
}

func (w *Wrapper[T]) Get(m *Metadata) (*T, error) {
	if m.DocumentType != w.docType {
		return nil, fmt.Errorf("wrong type")
	}
	var data T
	if err := m.Data.ToStruct(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
func (w *Wrapper[T]) Set(d *T) (*Metadata, error) {

	meta := &Metadata{
		DocumentType: w.docType,
		Version:      "0",
		Data:         make(JSONMap),
	}
	if err := meta.Data.FromStruct(d); err != nil {
		return nil, err
	}
	return meta, nil
}
