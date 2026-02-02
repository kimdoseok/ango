package main

type (
	Group struct {
		ID     int32      `gorm:"primary_key" json:"id"`
		Alumni []*Alumnus `gorm:"many2many:alumni_groups;"`
		Name   string     `gorm:"type:varchar(64); default:''; not null" json:"name"`
		Memo   string     `gorm:"type:text;" json:"memo"`
	}
)
