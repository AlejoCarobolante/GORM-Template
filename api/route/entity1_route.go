package route
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 9c06d0677de66b7a4a0df5652761fb7c46266c25

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
<<<<<<< HEAD
=======
>>>>>>> ab1ef724062c7782911371d261ff2370586a3ec1
=======
>>>>>>> 9c06d0677de66b7a4a0df5652761fb7c46266c25
