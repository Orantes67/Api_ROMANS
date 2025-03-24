package controllers

import (
	"ApiProduct/src/pagos/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListPagosController struct {
	useCase *application.ListPagos
}

func NewListPagosController(useCase *application.ListPagos) *ListPagosController {
	return &ListPagosController{useCase: useCase}
}

func (ctrl *ListPagosController) Execute(c *gin.Context) {
	pagos, err := ctrl.useCase.Execute() // Ahora solo un argumento
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pagos)
}
