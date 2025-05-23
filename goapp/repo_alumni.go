package main

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type (
	Alumni struct {
		ID             int    `gorm:"primary_key" json:"id"`
		FirstName      string `gorm:"type:varchar(64); default:''; not null" json:"first_name"`
		LastName       string `gorm:"type:varchar(64); default:''; not null" json:"last_name"`
		Title          string `gorm:"type:varchar(64); default:''; not null" json:"title"`
		Major          string `gorm:"type:varchar(32); default:''; not null" json:"major"`
		GraduationYear int    `gorm:"type:integer; default:0; not null" json:"graduation_year"`
		Email          string `gorm:"type:varchar(64); default:''; not null" json:"email"`
		Phone          string `gorm:"type:varchar(32); default:''; not null" json:"phone"`
		Address        string `gorm:"type:varchar(64); default:''; not null" json:"address"`
		City           string `gorm:"type:varchar(64); default:''; not null" json:"city"`
		State          string `gorm:"type:varchar(16); default:''; not null" json:"state"`
		Country        string `gorm:"type:varchar(32); default:''; not null" json:"country"`
		ZipCode        string `gorm:"type:varchar(16); default:''; not null" json:"zip_code"`
		Company        string `gorm:"type:varchar(64); default:''; not null" json:"company"`
		Position       string `gorm:"type:varchar(32); default:''; not null" json:"position"`
		WorkEmail      string `gorm:"type:varchar(64); default:''; not null" json:"work_email"`
		WorkPhone      string `gorm:"type:varchar(64); default:''; not null" json:"work_phone"`
		WorkAddress    string `gorm:"type:varchar(64); default:''; not null" json:"work_address"`
		WorkCity       string `gorm:"type:varchar(64); default:''; not null" json:"work_city"`
		WorkState      string `gorm:"type:varchar(16); default:''; not null" json:"work_state"`
		WorkCountry    string `gorm:"type:varchar(32); default:''; not null" json:"work_country"`
		Memo           string `gorm:"type:text;" json:"work_zip_code"`
	}

	AlumniRepository struct {
		Db *gorm.DB
	}
)

const (
	Limit int = 100
)

var (
	Offset int = 0
)

func (Alumni) TableName() string {
	return "alumni"
}

func NewAlumniRepository(db *gorm.DB) *AlumniRepository {
	//fmt.Println("NewErrorCodeRepository")

	return &AlumniRepository{
		Db: db,
	}
}

func (r *AlumniRepository) Get(ID int) (*Alumni, error) {
	var rec Alumni
	r.Db.Model(&Alumni{}).Where("id = ?", ID).First(&rec)
	//fmt.Println("Get Rec: =======>", code, rec)
	if &rec == nil || ID != rec.ID {
		return &Alumni{}, errors.New("Could not find the ID")
	}
	return &rec, nil
}

func (r *AlumniRepository) List(fstrs []string, page int) ([]Alumni, error) {
	var recs []Alumni
	//likestr := ""
	qstr, values := getConditionStr(fstrs)
	s := make([]interface{}, len(values))
	for i, v := range values {
		s[i] = v
	}
	Offset = page * Limit
	r.Db.Model(&Alumni{}).Where(qstr, s...).Order("code").Limit(Limit).Offset(Offset).Find(&recs)
	return recs, nil
}

func (r *AlumniRepository) ListAll(fstrs []string) ([]Alumni, error) {
	var recs []Alumni
	//likestr := ""
	qstr, values := getConditionStr(fstrs)
	s := make([]interface{}, len(values))
	for i, v := range values {
		s[i] = v
	}
	r.Db.Model(&Alumni{}).Where(qstr, s...).Order("code").Find(&recs)
	return recs, nil
}

func (r *AlumniRepository) Count(fstrs []string) int64 {
	var recs int64
	//likestr := ""
	qstr, values := getConditionStr(fstrs)
	s := make([]interface{}, len(values))
	for i, v := range values {
		s[i] = v
	}

	r.Db.Model(&Alumni{}).Where(qstr, s...).Count(&recs)
	return recs
}

func (r *AlumniRepository) Save(param *Alumni) (*Alumni, error) {
	var rec Alumni
	tx := r.Db.Begin()
	tx.Model(&Alumni{}).Where("code = ?", param.ID).Find(&rec)
	if rec.ID > 0 { // updatem
		tx.Model(&Alumni{}).Where("code = ?", rec.ID).Updates(param)
	} else { // create
		tx.Model(&Alumni{}).Create(param)
	}
	tx.Model(&Alumni{}).Where("code = ?", param.ID).First(&rec)
	tx.Commit()
	return &rec, nil
}

func (r *AlumniRepository) Delete(ID string) error {
	fmt.Println("ErrorCodeRepository Delete1", ID)
	tx := r.Db.Begin()
	tx.Where("code = ?", ID).Delete(&Alumni{})
	tx.Commit()
	return nil
}
