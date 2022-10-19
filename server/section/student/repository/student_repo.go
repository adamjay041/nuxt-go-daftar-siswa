package repository

import (
	"fmt"
	"server/helper"
	"server/model"

	"github.com/jinzhu/gorm"
)

type StudentRepo struct {
	DB *gorm.DB
}

func NewStudentRepo(DB *gorm.DB) model.StudentRepository {
	return &StudentRepo{DB}
}

func (m *StudentRepo) Fetch() (*[]model.Student, error) {
	var student []model.Student
	err := m.DB.Where("role = ?", "student").Find(&student).Error
	if err != nil {
		fmt.Printf("[PersonRepoImpl.ReadAll] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data")
	}
	return &student, nil
	// return &teacher, nil
}

func (m *StudentRepo) Create(t *model.Student) (*model.Student, error) {
	err := m.DB.Save(&t).Error
	if err != nil {
		fmt.Printf("[PersonRepoImpl.Create] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data")
	}
	return t, nil
}

func (m *StudentRepo) FetchBy(value string) (*[]model.Student, error) {
	// fmt.Printf("asdas %s", column)
	fmt.Print("ini value%", value)
	var student []model.Student
	err := m.DB.First(&student, "id = ?", value).Error
	if err != nil {
		// fmt.Printf("[PersonRepoImpl.ReadAll] error execute query %v \n", err)
		return nil, fmt.Errorf(err.Error())
	}
	return &student, nil
}

func (m *StudentRepo) EditBy(filter string, value string, t map[string]interface{}) (*model.Student, error) {
	var student model.Student
	// fmt.Printf(&t)
	if filter == "id" {
		err := m.DB.First(&student, "id = ?", value).Error
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}

		passTamp := &student.Password

		fmt.Printf("%v", t)

		if t["Password"] == "" {
			student.Password = *passTamp
		}

		t["Password"] = helper.DoHashUsingSalt(t["Password"].(string))

		errs := m.DB.First(&student, "id = ?", value).Update(t).Error
		if errs != nil {
			return nil, fmt.Errorf(errs.Error())
		}

	} else {
		err := m.DB.First(&student, "email = ?", value).Update(t).Error
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
	}

	return &student, nil

}

func (m *StudentRepo) DeleteBy(id string) (*model.Student, error) {
	var student model.Student
	err := m.DB.Delete(&student, id).Error
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return &student, nil
}
