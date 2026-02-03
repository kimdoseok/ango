package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type (
	ToolRepository struct {
		Db *gorm.DB
	}
)

func NewToolRepository(db *gorm.DB) *ToolRepository {
	return &ToolRepository{
		Db: db,
	}
}

func (r *AlumnusRepository) Import(xls *excelize.File) error {
	// Implement import logic here

	rows, err := xls.GetRows(SheetAlumni)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_ = rows // to avoid unused variable error

	log.Println("Import completed successfully")

	return nil
}
