package controller

import (
	"net/http"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Entity1Controller struct {
	Entity1Repository domain.Entity1Repository
}

func (te *Entity1Controller) Create(e *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var Entity1 domain.Entity1

	err := e.ShouldBind(&Entity1)
	if err != nil {
		e.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if Entity1.Name == "" {
		e.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Name is required"})
		return
	}

	Entity1.ID = uuid.New()

	err = te.Entity1Repository.Create(e, Entity1)
	if err != nil {
		e.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	e.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Entity1 created successfully",
	})
}

func (te *Entity1Controller) Fetch(e *gin.Context) {
	Entity1s, err := te.Entity1Repository.Fetch(e)
	if err != nil {
		e.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	e.JSON(http.StatusOK, Entity1s)
}

func (te *Entity1Controller) FetchById(e *gin.Context) {
	id := e.Param("id")

	Entity1s, err := te.Entity1Repository.FetchByID(e, id)
	if err != nil {
		e.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	e.JSON(http.StatusOK, Entity1s)
}

func (te *Entity1Controller) Update(e *gin.Context) {
	updatedEntity1 := &domain.Entity1{}

	err := e.ShouldBind(updatedEntity1)
	if err != nil {
		e.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedEntity1.ID == uuid.Nil {
		e.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID Entity1 is requiered to update"})
		return
	}

	err = te.Entity1Repository.Update(e, *updatedEntity1)
	if err != nil {
		e.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	e.JSON(http.StatusOK, domain.SuccessResponse{Message: "Entity1 updated succesfully"})
}

func (te *Entity1Controller) Delete(e *gin.Context) {
	Entity1ID := e.Param("id")
	err := te.Entity1Repository.Delete(e, Entity1ID)
	if err != nil {
		e.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	e.JSON(http.StatusOK, nil)
}
