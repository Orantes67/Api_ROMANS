package controllers

import (
	"ApiProduct/src/pagos/application"
	"ApiProduct/src/pagos/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)


type UpdatePagosController struct {
	useCase *application.UpdatePago
}

func NewUpdatePagosController(useCase *application.UpdatePago) *UpdatePagosController {
	return &UpdatePagosController{useCase: useCase}
}

func (ctrl *UpdatePagosController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var data struct {
		Descripcion string `json:"Descripcion" binding:"required"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPerfume := entities.NewPago(data.Descripcion)

	err = ctrl.useCase.Execute(id, updatedPerfume)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Perfume actualizado con éxito"})
}
