package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

func (this *Controller) test(c *gin.Context) {
	fmt.Println(this.service.dao.name)
	c.String(http.StatusOK, "hello word")
}

func (this *Controller) hello()  {
	fmt.Println("hello")
}

type Service struct {
	dao *Dao
}

func NewService(dao *Dao) *Service {
	return &Service{dao: dao}
}

type Dao struct {
	name string
}

func NewDao() *Dao {
	return &Dao{name: "数据库"}
}

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(NewDao)
	container.Provide(NewService)
	container.Provide(NewController)
	return container
}


func main()  {
	//监听端口默认为8080


	container:=BuildContainer()
	container.Invoke(func(controller *Controller) {
		r := gin.Default()
		r.GET("/", controller.test)
		r.Run(":8000")
		controller.hello()
	})
}