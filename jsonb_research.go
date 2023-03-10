package main

import (
	"fmt"
	"github.com/google/uuid"
	"swift-parser/model"
	"time"
)

func JsonbResearch() {

	repo := model.NewRepo()

	sukuk := model.Sukuk{
		WSID:   "0010",
		Param1: "111111",
	}

	sec := model.Security{
		ID:   uuid.New(),
		Name: fmt.Sprintf("security sukuk  %s ", time.Now().Format("2006-01-02 15:04:05")),
	}
	metadata, err := model.NewWrapper[model.Sukuk](model.DocTypeSukuk).Set(&sukuk)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sec.Metadata = metadata
	if err = repo.Create(&sec); err != nil {
		fmt.Println(err.Error())
		return
	}

	sec = model.Security{
		ID:   uuid.New(),
		Name: fmt.Sprintf("security external security %s ", time.Now().Format("2006-01-02 15:04:05")),
	}
	metadata, err = model.NewWrapper[model.ExternalSecurity](model.DocTypeExternalSecurity).
		Set(&model.ExternalSecurity{
			ISIN:   "0000003333",
			Param2: "param ----",
		})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sec.Metadata = metadata
	if err = repo.Create(&sec); err != nil {
		fmt.Println(err.Error())
		return
	}
	//securities, err := repo.GetAll()
	_, err = repo.GetAll()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//for _, row := range securities {
	//
	//	switch row.Metadata.DocumentType {
	//	case model.DocTypeSukuk:
	//		sukuk, err := model.NewWrapper[model.Sukuk](row.Metadata.DocumentType).Get(row.Metadata)
	//		if err != nil {
	//			fmt.Println(err.Error())
	//		} else {
	//			fmt.Printf("-- %s -- %s \n", sukuk.WSID, sukuk.Param1)
	//		}
	//
	//	case model.DocTypeExternalSecurity:
	//		es, err := model.NewWrapper[model.ExternalSecurity](row.Metadata.DocumentType).Get(row.Metadata)
	//		if err != nil {
	//			fmt.Println(err.Error())
	//		} else {
	//			fmt.Printf("-- %s -- %s \n", es.ISIN, es.Param2)
	//		}
	//
	//	}
	//
	//}

	sqlxRepo, err := model.NewSqlXRepo()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//sqlxSecurities, err := sqlxRepo.GetAll()
	_, err = sqlxRepo.GetAll()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//for _, row := range sqlxSecurities {
	//
	//	switch row.Metadata.DocumentType {
	//	case model.DocTypeSukuk:
	//		sukuk, err := model.NewWrapper[model.Sukuk](row.Metadata.DocumentType).Get(row.Metadata)
	//		if err != nil {
	//			fmt.Println(err.Error())
	//		} else {
	//			fmt.Printf("-- %s -- %s \n", sukuk.WSID, sukuk.Param1)
	//		}
	//
	//	case model.DocTypeExternalSecurity:
	//		es, err := model.NewWrapper[model.ExternalSecurity](row.Metadata.DocumentType).Get(row.Metadata)
	//		if err != nil {
	//			fmt.Println(err.Error())
	//		} else {
	//			fmt.Printf("-- %s -- %s \n", es.ISIN, es.Param2)
	//		}
	//
	//	}
	//
	//}

	_, err = repo.GetAllColumns()
	_, err = sqlxRepo.GetAllColumn()

}
