package controllers

import (
	"ApiProduct/src/perfumes_personalizados/application"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)

type DeletePerfumesPersonalizadoController struct{
	useCase *application.DeletePerfumePersonalizado
}

func NewDeletePerfumesPersonalizadoController(useCase *application.DeletePerfumePersonalizado) *DeletePerfumesPersonalizadoController{
	return &DeletePerfumesPersonalizadoController{useCase: useCase}
}
	func(ctrl *DeletePerfumesPersonalizadoController) Execute(c *gin.Context){
		idParam := c.Param("id")
		id,err := strconv.Atoi(idParam)
		if err !=nil {
			c.JSON(http.StatusBadRequest,gin.H{"error": "ID invalido"})
			return
		}
		err = ctrl.useCase.Execute(id)

		if err != nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"massage": "Perfume personalizado elimanado con éxito"})
	}
