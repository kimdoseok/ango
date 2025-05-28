package main

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type (
	Alumnus struct {
		ID             int      `gorm:"primary_key" json:"id"`
		Groups         []*Group `gorm:"many2many:alumni_groups;"`
		FirstName      string   `gorm:"type:varchar(64); default:''; not null" json:"first_name"`
		LastName       string   `gorm:"type:varchar(64); default:''; not null" json:"last_name"`
		Title          string   `gorm:"type:varchar(64); default:''; not null" json:"title"`
		Major          string   `gorm:"type:varchar(32); default:''; not null" json:"major"`
		GraduationYear int      `gorm:"type:integer; default:0; not null" json:"graduation_year"`
		Email          string   `gorm:"type:varchar(64); default:''; not null" json:"email"`
		Phone          string   `gorm:"type:varchar(32); default:''; not null" json:"phone"`
		Address        string   `gorm:"type:varchar(64); default:''; not null" json:"address"`
		City           string   `gorm:"type:varchar(64); default:''; not null" json:"city"`
		State          string   `gorm:"type:varchar(16); default:''; not null" json:"state"`
		Country        string   `gorm:"type:varchar(32); default:''; not null" json:"country"`
		ZipCode        string   `gorm:"type:varchar(16); default:''; not null" json:"zip_code"`
		Company        string   `gorm:"type:varchar(64); default:''; not null" json:"company"`
		Position       string   `gorm:"type:varchar(32); default:''; not null" json:"position"`
		WorkEmail      string   `gorm:"type:varchar(64); default:''; not null" json:"work_email"`
		WorkPhone      string   `gorm:"type:varchar(64); default:''; not null" json:"work_phone"`
		WorkAddress    string   `gorm:"type:varchar(64); default:''; not null" json:"work_address"`
		WorkCity       string   `gorm:"type:varchar(64); default:''; not null" json:"work_city"`
		WorkState      string   `gorm:"type:varchar(16); default:''; not null" json:"work_state"`
		WorkCountry    string   `gorm:"type:varchar(32); default:''; not null" json:"work_country"`
		Memo           string   `gorm:"type:text;" json:"memo"`
	}

	AlumnusRepository struct {
		Db *gorm.DB
	}
)

const (
	Limit int = 100
)

var (
	Offset int = 0
)

func (Alumnus) TableName() string {
	return "alumni"
}

func NewAlumnusRepository(db *gorm.DB) *AlumnusRepository {
	//fmt.Println("NewErrorCodeRepository")

	db.AutoMigrate(Alumnus{})
	return &AlumnusRepository{
		Db: db,
	}
}

func (r *AlumnusRepository) Get(ID int) (*Alumnus, error) {
	var rec Alumnus
	r.Db.Model(&Alumnus{}).Where("id = ?", ID).First(&rec)
	//fmt.Println("Get Rec: =======>", code, rec)
	if &rec == nil || ID != rec.ID {
		return &Alumnus{}, errors.New("Could not find the ID")
	}
	return &rec, nil
}

func (r *AlumnusRepository) List(fstrs []string, page int) ([]Alumnus, error) {
	var recs []Alumnus
	//likestr := ""
	qstr, values := getConditionStr(fstrs)
	s := make([]interface{}, len(values))
	for i, v := range values {
		s[i] = v
	}
	Offset = page * Limit
	r.Db.Model(&Alumnus{}).Where(qstr, s...).Order("code").Limit(Limit).Offset(Offset).Find(&recs)
	return recs, nil
}

func (r *AlumnusRepository) ListAll(fstrs []string) ([]Alumnus, error) {
	var recs []Alumnus
	//likestr := ""
	qstr, values := getConditionStr(fstrs)
	s := make([]interface{}, len(values))
	for i, v := range values {
		s[i] = v
	}
	r.Db.Model(&Alumnus{}).Where(qstr, s...).Order("code").Find(&recs)
	return recs, nil
}

func (r *AlumnusRepository) Count(fstrs []string) int64 {
	var recs int64
	//likestr := ""
	qstr, values := getConditionStr(fstrs)
	s := make([]interface{}, len(values))
	for i, v := range values {
		s[i] = v
	}

	r.Db.Model(&Alumnus{}).Where(qstr, s...).Count(&recs)
	return recs
}

func (r *AlumnusRepository) Save(param *Alumnus) (*Alumnus, error) {
	var rec Alumnus
	tx := r.Db.Begin()
	tx.Model(&Alumnus{}).Where("code = ?", param.ID).Find(&rec)
	if rec.ID > 0 { // updatem
		tx.Model(&Alumnus{}).Where("code = ?", rec.ID).Updates(param)
	} else { // create
		tx.Model(&Alumnus{}).Create(param)
	}
	tx.Model(&Alumnus{}).Where("code = ?", param.ID).First(&rec)
	tx.Commit()
	return &rec, nil
}

func (r *AlumnusRepository) Delete(ID string) error {
	fmt.Println("ErrorCodeRepository Delete1", ID)
	tx := r.Db.Begin()
	tx.Where("code = ?", ID).Delete(&Alumnus{})
	tx.Commit()
	return nil
}
