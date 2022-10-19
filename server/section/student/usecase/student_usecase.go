package usecase

import "server/model"

type StudentUseCase struct {
	studentRepo model.StudentRepository
}

func NewStudentUseCase(t model.StudentRepository) model.StudentUsecase {
	return &StudentUseCase{studentRepo: t}
}

func (m *StudentUseCase) Fecth() (*[]model.Student, error) {
	return m.studentRepo.Fetch()
}

func (m *StudentUseCase) Create(Student *model.Student) (*model.Student, error) {
	return m.studentRepo.Create(Student)
}

func (m *StudentUseCase) FetchBy(value string) (*[]model.Student, error) {
	return m.studentRepo.FetchBy(value)
}

func (m *StudentUseCase) EditBy(filter string, value string, t map[string]interface{}) (*model.Student, error) {
	return m.studentRepo.EditBy(filter, value, t)
}

func (m *StudentUseCase) DeleteBy(id string) (*model.Student, error) {
	return m.studentRepo.DeleteBy(id)
}
