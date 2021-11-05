package model


type Student struct {
	Id int `json:"id"`
	Name string `json:"name"`
	TeacherId int `json:"teacherId"`
	Teacher Teacher `gorm:"FOREIGNKEY:TeacherId;ASSOCIATION_FOREIGNKEY:ID"`
}