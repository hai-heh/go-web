package route

import (
	"github.com/gin-gonic/gin"
	"web/poject/controller"
)

type Route struct {
	c controller.StudentController
	engine *gin.Engine
}

func NewRoute(engine *gin.Engine) *Route {
	return &Route{c: controller.NewStudentControllerImpl(),engine:engine}
}

func (this *Route) AddStudentRoute() *Route {
	this.engine.GET("/",this.c.FindAll)
	this.engine.GET("/:id",this.c.FindById)
	this.engine.DELETE("/:id",this.c.DeleteById)
	this.engine.PUT("/:id",this.c.UpdateById)
	this.engine.POST("/",this.c.Create)
	return this
}
