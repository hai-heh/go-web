package dao

import (
	"web/poject/model"
)

type StudentDao interface {
	FindAll() ([]model.Student,error)
	FindById(id int) (model.Student,error)
	Create(student model.Student) error
	UpdateById(student model.Student) error
	DeleteById(id int) error
}

func NewUserDaoImpl() StudentDao {
	return &StudentDaoImpl{}
}


type StudentDaoImpl struct{

}

func (StudentDaoImpl) FindAll() ([]model.Student, error) {
	coon:=getCoon()
	var students []model.Student
	err:=coon.Find(&students).Error
	close(coon)
	return students,err

}

func (StudentDaoImpl) FindById(id int) (model.Student, error) {
	coon:=getCoon()
	var student model.Student
	err:=coon.First(&student,id).Error
	close(coon)
	return student,err
}

func (StudentDaoImpl) Create(student model.Student) error {
	coon:=getCoon()
	err:=coon.Create(&student).Error
	close(coon)
	return err
}

func (StudentDaoImpl) UpdateById(student model.Student) error {
	coon:=getCoon()
	err:=coon.Model(&student).Updates(map[string]interface{}{"name":student.Name,"teacherId":student.TeacherId}).Error
	close(coon)
	return err
}

func (StudentDaoImpl) DeleteById(id int) error {
	student:=model.Student{Id:id}
	coon:=getCoon()
	err:=coon.Delete(&student).Error
	close(coon)
	return err
}






