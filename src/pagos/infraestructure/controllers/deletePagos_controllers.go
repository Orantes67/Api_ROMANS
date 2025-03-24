package controllers

import (
	"ApiProduct/src/pagos/application"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)

type DeletePagosController struct {
	useCase *application.DeletePagos
}

func NewDeletePagosController(useCase *application.DeletePagos) *DeletePagosController {
	return &DeletePagosController{useCase: useCase}
}

	func (ctrl *DeletePagosController) Execute(c *gin.Context){
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}
	
		err = ctrl.useCase.Execute(id) // Ahora pasamos el ID correctamente
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{"message": "Pago eliminado con éxito"})
	}
	
