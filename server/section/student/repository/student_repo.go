package repository

import (
	"fmt"
	"log"
	"server/dto"
	"server/helper"
	"server/model"
	"strconv"
	"strings"

	"github.com/mashingan/smapping"
	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
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

func (m *StudentRepo) Create(t *dto.CreateRequestStudent) (*dto.CreateRequestStudent, error) {

	student := model.Student{}

	er := smapping.FillStruct(&student, smapping.MapFields(&t))

	if er != nil {
		log.Fatalf("Failed map %v", er)
		return nil, er
	}

	tampArr := strings.Split(student.Periode, "-")
	intVar, er := strconv.Atoi(tampArr[0])
	student.Status = "active"

	if er != nil {
		fmt.Printf("Failed To Generate Generation")
	}

	if intVar%2 == 0 {
		student.Generation = fmt.Sprintf("%d/%d Genap", intVar, intVar+1)
	} else {
		student.Generation = fmt.Sprintf("%d/%d Ganjil", intVar, intVar+1)
	}

	err := m.DB.Save(&student).Error
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
		return nil, fmt.Errorf(err.Error())
	}
	return &student, nil
}

func (m *StudentRepo) EditBy(filter string, value string, t *dto.EditRequestStudent) (*model.Student, error) {
	student := model.Student{}

	er := smapping.FillStruct(&student, smapping.MapFields(&t))

	if er != nil {
		log.Fatalf("Failed map %v", er)
		return nil, er
	}
	fmt.Println(student)

	if filter == "id" {
		err := m.DB.First(&model.Student{}, "id = ?", value).Error
		if err != nil {
			// fmt.Printf("[PersonRepoImpl.ReadAll] error execute query %v \n", err)

			return nil, fmt.Errorf(err.Error())
		}

		passTamp := &student.Password

		// fmt.Printf("%v", passTamp)

		if student.Password == "" {
			student.Password = *passTamp
		}

		student.Password = helper.DoHashUsingSalt(student.Password)

		errs := m.DB.First(&model.Student{}, "id = ?", value).Updates(student).Error
		// errs := m.DB.Model(&model.Student{}).Where("id = ?", value).Save(student).Error
		if errs != nil {
			fmt.Println("MASUK SINI KONTOL")
			return nil, fmt.Errorf(errs.Error())
		}

	} else {
		err := m.DB.First(&student, "email = ?", value).Updates(t).Error
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
	}

	return &student, nil

}

func (m *StudentRepo) DeleteBy(id string) (*model.Student, error) {
	var student model.Student
	err := m.DB.Where("id = ?", id).Delete(&model.Student{}).Error
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return &student, nil
}
