package main

import (
	"errors"

	"gorm.io/gorm"
)

type (

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

func (r *GroupRepository) Get(ID int32) (*Group, error) {
	var rec Group
	r.Db.Model(&Group{}).Where("id = ?", ID).First(&rec)
	//fmt.Println("Get Rec: =======>", code, rec)
	if &rec == nil || ID == rec.ID {
		return &Group{}, errors.New("Could not find the ID")
	}
	return &rec, nil
}

func (r *GroupRepository) List(fstrs []string, page int) ([]Group, error) {
	var recs []Group
	fstr := ""
	if len(fstrs) > 0 {
		fstr = fstrs[0]
	}
	GroupOffset = page * GroupLimit
	//r.Db.Model(&Group{}).Joins("left outer join alumni_groups on alumni_groups.group_id = groups.id").Where("groups.name like '%%%s%%' ", fstr).Order("name").Limit(GroupLimit).Offset(GroupOffset).Find(&recs)
	r.Db.Model(&Group{}).Where("groups.name like '%%%s%%' ", fstr).Order("name").Limit(GroupLimit).Offset(GroupOffset).Find(&recs)
	return recs, nil
}

func (r *GroupRepository) ListAll(fstrs []string) ([]Group, error) {
	var recs []Group

	r.Db.Model(&Group{}).Order("name").Limit(GroupLimit).Offset(GroupOffset).Find(&recs)
	return recs, nil
}

func (r *GroupRepository) Count(fstrs []string) int64 {
	var num int64
	//likestr := ""

	fstr := ""
	if len(fstrs) > 0 {
		fstr = fstrs[0]
	}
	r.Db.Model(&Group{}).Where("groups.name like '%%%s%%' ", fstr).Count(&num)
	return num
}

func (r *GroupRepository) Save(param *Group) (*Group, error) {
	var rec Group
	tx := r.Db.Begin()
	tx.Model(&Group{}).Where("id = ?", param.ID).Find(&rec)
	if rec.ID > 0 { // updatem
		tx.Model(&Group{}).Where("id = ?", rec.ID).Updates(param)
	} else { // create
		tx.Model(&Group{}).Create(param)
	}
	tx.Model(&Group{}).Where("id = ?", param.ID).First(&rec)
	tx.Commit()
	return &rec, nil
}

func (r *GroupRepository) Delete(ID int32) error {
	tx := r.Db.Begin()
	tx.Where("id = ?", ID).Delete(&Group{})
	tx.Commit()
	return nil
}
