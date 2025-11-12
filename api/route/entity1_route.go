package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewEntityRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.Entity1Controller{
		Entity1Repository: &usecase.Entity1UseCase{},
	}
	Entity1Router := group.Group("/entity1")
	Entity1Router.POST("/", ec.Create)
	Entity1Router.GET("/", ec.Fetch)
	Entity1Router.GET("/:id", ec.FetchById)
	Entity1Router.PUT("/", ec.Update)
	Entity1Router.DELETE("/:id", ec.Delete)
}