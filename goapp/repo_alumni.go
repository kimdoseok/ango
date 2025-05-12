package main

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type (
	Alumni struct {
		ID             int    `json:"id"`
		FirstName      string `json:"first_name"`
		LastName       string `json:"last_name"`
		Title          string `json:"title"`
		Major          string `json:"major"`
		GraduationYear int    `json:"graduation_year"`
		Email          string `json:"email"`
		Phone          string `json:"phone"`
		Address        string `json:"address"`
		City           string `json:"city"`
		State          string `json:"state"`
		Country        string `json:"country"`
		ZipCode        string `json:"zip_code"`
		Company        string `json:"company"`
		Position       string `json:"position"`
		WorkEmail      string `json:"work_email"`
		WorkPhone      string `json:"work_phone"`
		WorkAddress    string `json:"work_address"`
		WorkCity       string `json:"work_city"`
		WorkState      string `json:"work_state"`
		WorkCountry    string `json:"work_country"`
		WorkZipCode    string `json:"work_zip_code"`
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
	db.AutoMigrate(Alumni{})

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
