package main

import (
	"errors"

	//"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type (
	AlumnusRepository struct {
		Db *gorm.DB
	}
)

const (
	AlumnusLimit int    = 100
	SheetAlumni  string = "Alumni"
	SheetGroups  string = "Groups"
/*
	fieldmap [][]string = [][]string{
		[]string{"ID", "s", "A"},
		[]string{"FirstName", "s", "B"},
		[]string{"LastName", "s", "C"},
		[]string{"Title", "s", "D"},
		[]string{"Major", "s", "E"},
		[]string{"GraduationYear", "i", "F"},
		[]string{"Email", "s", "G"},
		[]string{"Phone", "s", "H"},
		[]string{"Address", "s", "I"},
		[]string{"City", "s", "J"},
		[]string{"State", "s", "K"},
		[]string{"Country", "s", "L"},
		[]string{"ZipCode", "s", "M"},
		[]string{"Company", "s", "N"},
		[]string{"Position", "s", "O"},
		[]string{"WorkEmail", "s", "P"},
		[]string{"WorkPhone", "s", "Q"},
		[]string{"WorkAddress", "s", "R"},
		[]string{"WorkCity", "s", "S"},
		[]string{"WorkState", "s", "T"},
		[]string{"WorkCountry", "s", "U"},
		[]string{"Memo", "s", "V"},
	}
*/
)


var (
	AlumnusOffset int = 0
	Columns       []string
	//Columns       []string = {"ID", "FirstName", "LastName", "Title", "Major", "GraduationYear", "Email", "Phone", "Address", "City", "State", "Country", "ZipCode", "Company", "Position", "WorkEmail", "WorkPhone", "WorkAddress", "WorkCity", "WorkState", "WorkCountry", "Memo"}
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

func (r *AlumnusRepository) Get(ID int32) (*Alumnus, error) {
	var rec Alumnus
	r.Db.Model(&Alumnus{}).Where("id = ?", ID).First(&rec)
	//fmt.Println("Get Rec: =======>", code, rec)
	if &rec == nil || ID != int32(rec.ID) {
		return &Alumnus{}, errors.New("Could not find the ID")
	}
	return &rec, nil
}

func (r *AlumnusRepository) List(fstrs []string, page int) ([]Alumnus, error) {
	var recs []Alumnus
	fstr := ""
	if len(fstrs) > 0 {
		fstr = fstrs[0]
	}
	AlumnusOffset = page * AlumnusLimit

	//r.Db.Model(&Alumnus{}).Joins("left outer join alumni_Alumnuss on alumni_Alumnuss.Alumnus_id = Alumnuss.id").Where("Alumnuss.name like '%%%s%%' ", fstr).Order("name").Limit(AlumnusLimit).Offset(AlumnusOffset).Find(&recs)
	r.Db.Model(&Alumnus{}).Where("Alumnuss.first_name like '%%%s%%' OR Alumnuss.last_name like '%%%s%%' ", fstr).Order("last_name, first_name").Limit(AlumnusLimit).Offset(AlumnusOffset).Find(&recs)
	return recs, nil
}

func (r *AlumnusRepository) ListAll(fstrs []string) ([]Alumnus, error) {
	var recs []Alumnus

	r.Db.Model(&Alumnus{}).Order("last_name, first_name").Find(&recs)
	return recs, nil
}

func (r *AlumnusRepository) Count(fstrs []string) int64 {
	var num int64
	//likestr := ""

	fstr := ""
	if len(fstrs) > 0 {
		fstr = fstrs[0]
	}
	r.Db.Model(&Alumnus{}).Where("alumnuss.name like '%%%s%%' ", fstr).Count(&num)
	return num
}

func (r *AlumnusRepository) Save(param *Alumnus) (*Alumnus, error) {
	var rec Alumnus
	tx := r.Db.Begin()
	tx.Model(&Alumnus{}).Where("id = ?", param.ID).Find(&rec)
	if rec.ID > 0 { // updatem
		tx.Model(&Alumnus{}).Where("id = ?", rec.ID).Updates(param)
	} else { // create
		tx.Model(&Alumnus{}).Create(param)
	}
	tx.Model(&Alumnus{}).Where("id = ?", param.ID).First(&rec)
	tx.Commit()
	return &rec, nil
}

func (r *AlumnusRepository) Delete(ID int32) error {
	tx := r.Db.Begin()
	tx.Where("id = ?", ID).Delete(&Alumnus{})
	tx.Commit()
	return nil
}
