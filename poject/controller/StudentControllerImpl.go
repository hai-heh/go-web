package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/poject/model"
	"web/poject/service"
)

type StudentController interface {
	 FindAll(gin *gin.Context)
	 FindById(ctx *gin.Context)
	 Create(ctx *gin.Context)
	 DeleteById(ctx *gin.Context)
	 UpdateById(ctx *gin.Context)
}



func NewStudentControllerImpl() StudentController {
	return &StudentControllerImpl{studentService:service.NewUserServiceImple()}
}


type StudentControllerImpl struct {
	studentService service.StudentService

}


func (this *StudentControllerImpl) FindAll(ctx *gin.Context)  {
	students, e := this.studentService.FindAll()
	if(e!=nil){
		ctx.JSON(500,gin.H{"errors":e})
		return
	}
	ctx.JSON(200,gin.H{"students":students})
}

func (this *StudentControllerImpl) FindById(ctx *gin.Context) {
	 id,err:=strconv.Atoi(ctx.Param("id"))
	 if(err!=nil){
	 	ctx.JSON(400,gin.H{"erros":err})
		 return
	 }
	if student, err := this.studentService.FindById(id);err!=nil{
		ctx.JSON(500,gin.H{"erros":err})
		return
	}else{
		ctx.JSON(200,gin.H{"student":student})
	}
}


func (this *StudentControllerImpl) Create(ctx *gin.Context) {
	var student model.Student
    if err:=ctx.ShouldBindJSON(&student);err!=nil{
    	ctx.JSON(400,gin.H{"erros":err})
		return
	}
	if err:=this.studentService.Create(student);err!=nil{
		ctx.JSON(500,gin.H{"erros":err})
		return
	}
	ctx.JSON(200,"添加成功")
}

func (this *StudentControllerImpl) DeleteById(ctx *gin.Context) {
	id,err:=strconv.Atoi(ctx.Param("id"))
	if(err!=nil){
		ctx.JSON(400,gin.H{"erros":err})
		return
	}
	if err := this.studentService.DeleteById(id);err!=nil{
		ctx.JSON(500,gin.H{"erros":err})
		return
	}
	ctx.JSON(200,"删除成功")
}

func (this *StudentControllerImpl) UpdateById(ctx *gin.Context) {
	var student model.Student
	if err:=ctx.ShouldBindJSON(&student);err!=nil{
		ctx.JSON(400,gin.H{"erros":err})
		return
	}
	if err:=this.studentService.UpdateById(student);err!=nil{
		ctx.JSON(500,gin.H{"erros":err})
		return
	}
	ctx.JSON(200,"修改成功")

}




