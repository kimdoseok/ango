package main

import (
	"errors"

	"gorm.io/gorm"
)

type (
	Group struct {
		ID     int        `gorm:"primary_key" json:"id"`
		Alumni []*Alumnus `gorm:"many2many:alumni_groups;"`
		Name   string     `gorm:"type:varchar(64); default:''; not null" json:"name"`
		Memo   string     `gorm:"type:text;" json:"memo"`
	}

	GroupRepository struct {
		Db *gorm.DB
	}
)

const (
	GroupLimit int = 100
)

var (
	GroupOffset int = 0
)

func (Group) TableName() string {
	return "groups"
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	//fmt.Println("NewErrorCodeRepository")

	db.AutoMigrate(Group{})
	return &GroupRepository{
		Db: db,
	}
}

func (r *GroupRepository) Get(ID int) (*Group, error) {
	var rec Group
	r.Db.Model(&Group{}).Where("id = ?", ID).First(&rec)
	//fmt.Println("Get Rec: =======>", code, rec)
	if &rec == nil || ID != rec.ID {
		return &Group{}, errors.New("Could not find the ID")
	}
	return &rec, nil
}

func (r *GroupRepository) List(fstrs []string, page int) ([]Group, error) {
	var recs []Group
	//likestr := ""
	qstr, values := getConditionStr(fstrs)
	s := make([]interface{}, len(values))
	for i, v := range values {
		s[i] = v
	}
	GroupOffset = page * GroupLimit
	r.Db.Model(&Group{}).Where(qstr, s...).Order("code").Limit(GroupLimit).Offset(GroupOffset).Find(&recs)
	return recs, nil
}

func (r *GroupRepository) ListAll(fstrs []string) ([]Group, error) {
	var recs []Group
	//likestr := ""
	qstr, values := getConditionStr(fstrs)
	s := make([]interface{}, len(values))
	for i, v := range values {
		s[i] = v
	}
	r.Db.Model(&Group{}).Where(qstr, s...).Order("code").Find(&recs)
	return recs, nil
}

func (r *GroupRepository) Count(fstrs []string) int64 {
	var num int64
	//likestr := ""
	qstr, values := getConditionStr(fstrs)
	s := make([]interface{}, len(values))
	for i, v := range values {
		s[i] = v
	}

	r.Db.Model(&Group{}).Where(qstr, s...).Count(&num)
	return num
}

func (r *GroupRepository) Save(param *Group) (*Group, error) {
	var rec Group
	tx := r.Db.Begin()
	tx.Model(&Group{}).Where("code = ?", param.ID).Find(&rec)
	if rec.ID > 0 { // updatem
		tx.Model(&Group{}).Where("code = ?", rec.ID).Updates(param)
	} else { // create
		tx.Model(&Group{}).Create(param)
	}
	tx.Model(&Group{}).Where("code = ?", param.ID).First(&rec)
	tx.Commit()
	return &rec, nil
}

func (r *GroupRepository) Delete(ID string) error {
	tx := r.Db.Begin()
	tx.Where("code = ?", ID).Delete(&Group{})
	tx.Commit()
	return nil
}
