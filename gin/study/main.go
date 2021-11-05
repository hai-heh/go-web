package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type FooHdl struct {

}


type barForm struct {
	Foo string  `form:"foo" binding:"required"`
	Bar int     `form:"bar" binding:"required"`
}

func (fooHdl *FooHdl) Bar(c *gin.Context) {
	var bform = new(barForm)
	if err := c.ShouldBind(bform); err != nil {
		// true: parse form error
		return
	}

	// handle biz logic and generate response structure
	// c (gin.Context) methods could called to support process-controling

	c.JSON(http.StatusOK, "成功")
	// c.String() alse repsonse to client
}

// mountRouters .
func mountRouters(engi *gin.Engine) {
	// use middlewares
	engi.Use(gin.Logger())
	engi.Use(gin.Recovery())
	engi.GET("redis",(new (FooHdl)).Bar)
	engi.GET("redi",(new (FooHdl)).Bar)
	// mount routers

	group := engi.Group("/v1")
	{

		fooHdl := FooHdl{}
		group.GET("/foo", fooHdl.Bar)
		//group.GET("/echo", fooHdl.Echo)
		// subGroup := group.Group("/subg")
		// subGroup.GET("/hdl1", fooHdl.SubGroupHdl1) // 最终路由："targetURI = /v1/subg/hdl1"
	}
}

func main() {
	engi := gin.New()


	mountRouters(engi)

	if err := engi.Run(":8080"); err != nil {
		log.Fatalf("engi exit with err=%v", err)
	}
}
