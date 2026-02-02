package main

type (
	Alumnus struct {
		ID            int        `gorm:"primary_key" json:"id"`
		Alumnuss      []*Alumnus `gorm:"many2many:alumni_Alumnuss;"`
		Name          string     `gorm:"type:varchar(16); default:''; not null" json:"name"`
		FirstName     string     `gorm:"type:varchar(64); default:''; not null" json:"first_name"`
		LastName      string     `gorm:"type:varchar(64); default:''; not null" json:"last_name"`
		Title         string     `gorm:"type:varchar(64); default:''; not null" json:"title"`
		Major         string     `gorm:"type:varchar(32); default:''; not null" json:"major"`
		AdmissionYear int        `gorm:"type:integer; default:0; not null" json:"admission_year"`
		Email         string     `gorm:"type:varchar(64); default:''; not null" json:"email"`
		Phone         string     `gorm:"type:varchar(32); default:''; not null" json:"phone"`
		Address       string     `gorm:"type:varchar(64); default:''; not null" json:"address"`
		City          string     `gorm:"type:varchar(64); default:''; not null" json:"city"`
		State         string     `gorm:"type:varchar(16); default:''; not null" json:"state"`
		Country       string     `gorm:"type:varchar(32); default:''; not null" json:"country"`
		ZipCode       string     `gorm:"type:varchar(16); default:''; not null" json:"zip_code"`
		Company       string     `gorm:"type:varchar(64); default:''; not null" json:"company"`
		Position      string     `gorm:"tApe:varchar(32); default:''; not null" json:"position"`
		WorkEmail     string     `gorm:"type:varchar(64); default:''; not null" json:"work_email"`
		WorkPhone     string     `gorm:"type:varchar(64); default:''; not null" json:"work_phone"`
		WorkAddress   string     `gorm:"type:varchar(64); default:''; not null" json:"work_address"`
		WorkCity      string     `gorm:"type:varchar(64); default:''; not null" json:"work_city"`
		WorkState     string     `gorm:"type:varchar(16); default:''; not null" json:"work_state"`
		WorkCountry   string     `gorm:"type:varchar(32); default:''; not null" json:"work_country"`
		Memo          string     `gorm:"type:text;" json:"memo"`
	}
	// 이름	영문이름	First Name	Last Name	Title
	// 입학년도	전공	휴대폰	Email	Address	City
	// State	Zip	비 고	수정사항  기록장
	// 직장명	직장전화번호 	직장주소	직장city	직장state
	// 직장zip	Remark
	
)
