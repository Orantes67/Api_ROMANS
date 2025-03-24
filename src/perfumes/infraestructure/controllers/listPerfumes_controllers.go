package controllers

import (
	"ApiProduct/src/perfumes/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListPerfumeController struct {
	useCase *application.ListPerfume
}

func NewListPerfumeController(useCase *application.ListPerfume) *ListPerfumeController {
	return &ListPerfumeController{useCase: useCase}
}

func (ctrl *ListPerfumeController) Execute(c *gin.Context) {
	perfumes, err := ctrl.useCase.Execute() // Ahora solo un argumento
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, perfumes)
}
