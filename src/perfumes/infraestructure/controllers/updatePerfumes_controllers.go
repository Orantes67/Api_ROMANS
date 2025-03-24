package controllers

import (
	"ApiProduct/src/perfumes/application"
	"ApiProduct/src/perfumes/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)


type UpdatePerfumeController struct {
	useCase *application.UpdatePerfume
}

func NewUpdatePerfumeController(useCase *application.UpdatePerfume) *UpdatePerfumeController {
	return &UpdatePerfumeController{useCase: useCase}
}

func (ctrl *UpdatePerfumeController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var data struct {
		Nombre    string `json:"nombre" binding:"required"`
		Descripcion string `json:"Descripcion" binding:"required"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPerfume := entities.NewPerfume(data.Nombre,data.Descripcion)

	err = ctrl.useCase.Execute(id, updatedPerfume)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Perfume actualizado con éxito"})
}
