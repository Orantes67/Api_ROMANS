package controllers

import (
	"ApiProduct/src/pagos/application"
	"ApiProduct/src/pagos/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)


type CreatePagosController struct{
	useCase *application.CreatePago
}

func NewCreatePagosController(useCase *application.CreatePago) *CreatePagosController{
	return &CreatePagosController{useCase: useCase}
}

func (cpay_c *CreatePagosController) Execute(c *gin.Context){
	var data struct{
		Descripcion string `json:"Descripcion" binding:"required"`
	}

	if err := c.ShouldBindJSON(&data); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pagos := entities.NewPago(data.Descripcion)

	err := cpay_c.useCase.Execute(pagos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated,gin.H{"message": "Pago creado con éxito"})
}