package main

import "gorm.io/gorm"

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

	rows, err := f.GetRows(SheetAlumni)
	if err != nil {
		fmt.Println(err)
		return f
	}

	for i, row := range rows {
		lenrow := len(row)
		fmt.Println(row.Columns())

		if i < 1 {
			continue
		}

		if lenrow < 3 {
			continue
		}
		if len(row[1]) < 1 {
			continue
		}
		a := &Alumnus{}
		j := 1
		cellname, err := excelize.CoordinatesToCellName(j, i+1)
		if err != nil {
			log.Println(err)
		}
		cellstr, err := xls.GetCellValue(SheetAlumni, cellname)
		if err != nil {
			log.Println(err)
		}
		a.FirstName = cellstr
		j += 1
		cellname, err = excelize.CoordinatesToCellName(j, i+1)
		if err != nil {
			log.Println(err)
		}
		cellstr, err = xls.GetCellValue(SheetAlumni, cellname)
		if err != nil {
			log.Println(err)
		}
		a.LastName = cellstr

		j += 1
		cellname, err = excelize.CoordinatesToCellName(j, i+1)
		if err != nil {
			log.Println(err)
		}
		cellstr, err = xls.GetCellValue(SheetAlumni, cellname)
		if err != nil {
			log.Println(err)
		}
		a.Title = cellstr


		
		sv := func(a *Alumnus, row int) {
			
			for i, fld := range fieldmap {
				cellname, err := excelize.CoordinatesToCellName(i, row+1)
				if err != nil {
					log.Println(err)
				}
				cellstr, err := xls.GetCellValue(SheetAlumni, cellname)
				if err != nil {
					log.Println(err)
				}

				if fld[0] == col {
					switch fld[1] {
					case "s":
						// Set string value
						//reflect.ValueOf(a).Elem().FieldByName(fld[0]).SetString(value)
						reflect.ValueOf(a).Elem().FieldByName(fld[0]).SetString(cellstr)
					case "i":
						// Set integer value
						var intValue int
						fmt.Sscanf(cellstr, "%d", &intValue)
						reflect.ValueOf(a).Elem().FieldByName(fld[0]).SetInt(int64(intValue))
					}
					break
				}
				fmt.Println("Field: ", fld[0], " Value: ", cellstr)
			}

			t := reflect.TypeOf(a)
			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				if field.Name == name {
					field.Value = value
					break
				}
			}
			t := reflect.TypeOf(a)
			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				fmt.Println(field.Name)
			j += 1
			cellname, err = excelize.CoordinatesToCellName(j, i+1)
			if err != nil {
				log.Println(err)
			}
			cellstr, err = xls.GetCellValue(SheetAlumni, cellname)
			if err != nil {
				log.Println(err)
			}
			switch col {
			case "AdmissionYear":
				fmt.Sscanf(cellstr, "%d", &s.AdmissionYear)
			case "Email":
				s.Email = cellstr
			case "Phone":
				s.Phone = cellstr
		/*
			a.ID =              int        `gorm:"primary_key" json:"id"`

			Alumnuss       []*Alumnus `gorm:"many2many:alumni_Alumnuss;"`
			FirstName      string     `gorm:"type:varchar(64); default:''; not null" json:"first_name"`
			LastName       string     `gorm:"type:varchar(64); default:''; not null" json:"last_name"`
			Title          string     `gorm:"type:varchar(64); default:''; not null" json:"title"`
			Major          string     `gorm:"type:varchar(32); default:''; not null" json:"major"`
			GraduationYear int        `gorm:"type:integer; default:0; not null" json:"graduation_year"`
			Email          string     `gorm:"type:varchar(64); default:''; not null" json:"email"`
			Phone          string     `gorm:"type:varchar(32); default:''; not null" json:"phone"`
			Address        string     `gorm:"type:varchar(64); default:''; not null" json:"address"`
			City           string     `gorm:"type:varchar(64); default:''; not null" json:"city"`
			State          string     `gorm:"type:varchar(16); default:''; not null" json:"state"`
			Country        string     `gorm:"type:varchar(32); default:''; not null" json:"country"`
			ZipCode        string     `gorm:"type:varchar(16); default:''; not null" json:"zip_code"`
			Company        string     `gorm:"type:varchar(64); default:''; not null" json:"company"`
			Position       string     `gorm:"tApe:varchar(32); default:''; not null" json:"position"`
			WorkEmail      string     `gorm:"type:varchar(64); default:''; not null" json:"work_email"`
			WorkPhone      string     `gorm:"type:varchar(64); default:''; not null" json:"work_phone"`
			WorkAddress    string     `gorm:"type:varchar(64); default:''; not null" json:"work_address"`
			WorkCity       string     `gorm:"type:varchar(64); default:''; not null" json:"work_city"`
			WorkState      string     `gorm:"type:varchar(16); default:''; not null" json:"work_state"`
			WorkCountry    string     `gorm:"type:varchar(32); default:''; not null" json:"work_country"`
			Memo           string     `gorm:"type:text;" json:"memo"`
		*/

		if i < 2 {
			continue
		}




		j := 1
		cellname, err = excelize.CoordinatesToCellName(j, i+1)
		if err != nil {
			log.Println(err)
		}
		cellstr, err = xls.GetCellValue(SheetAlumni, cellname)
		if err != nil {
			log.Println(err)
		}
		a.Major = cellstr
		
		s := reflect.ValueOf(a).Elem()
		field := s.FieldByName("Field1")
		if field.IsValid() && field.CanSet() {
			field.SetString("New Value") // For a string field
			field.SetInt(123)
		}

		f.SetCellValue(sheetname, cellname, last)
		cellname, err = excelize.JoinCellName("E", i+1)
		if err != nil {
			continue
		}
		f.SetCellValue(sheetname, cellname, title)
		//fmt.Println(SplitName(row[1]))
		fmt.Println(first)
		fmt.Println(last)
		fmt.Println(title)
	}
	f.Save()

	return nil
}
