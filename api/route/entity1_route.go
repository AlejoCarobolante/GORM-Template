package route

import (
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewEntityRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.Entity1Controller{
		Entity1Repository: &usecase.Entity1UseCase{},
	}
	entity1Routes := group.Group("/entity1")
	entity1Routes.POST("", ec.Create)
	entity1Routes.GET("", ec.Fetch)
	entity1Routes.GET(":id", ec.FetchById)
	entity1Routes.PUT("", ec.Update)
	entity1Routes.DELETE(":id", ec.Delete)
}
