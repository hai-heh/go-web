package service

import (
	"web/poject/dao"
	"web/poject/model"
)

type StudentService interface {
	FindAll() ([]model.Student,error)
	FindById(id int) (model.Student,error)
	Create(student model.Student) error
	UpdateById(student model.Student) error
	DeleteById(id int) error
}

type StudentServiceImpl struct{
	StudentDao dao.StudentDao
}

func NewUserServiceImple() *StudentServiceImpl {
	return &StudentServiceImpl{StudentDao:dao.NewUserDaoImpl()}
}

func (this *StudentServiceImpl) FindAll() ([]model.Student,error) {
	return this.StudentDao.FindAll()
}

func (this *StudentServiceImpl) FindById(id int) (model.Student,error) {
	return this.StudentDao.FindById(id)
}

func (this *StudentServiceImpl) Create(student model.Student) error {
	return this.StudentDao.Create(student)
}

func (this *StudentServiceImpl) UpdateById(student model.Student) error {
	return this.StudentDao.UpdateById(student)
}

func (this *StudentServiceImpl) DeleteById(id int) error {
	return this.StudentDao.DeleteById(id)
}

