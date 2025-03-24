package controllers

import (
	"ApiProduct/src/perfumes_personalizados/application"
	"ApiProduct/src/perfumes_personalizados/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePerfumePersonalizadoController struct{
	useCase *application.CreatePerfumePersonalizado
}

func NewCreatePefumesPersonalizadoController(useCase *application.CreatePerfumePersonalizado) *CreatePerfumePersonalizadoController{
	return &CreatePerfumePersonalizadoController{useCase: useCase}
}

func (cpp_c *CreatePerfumePersonalizadoController) Execute(c *gin.Context){
	var data struct{
		Nombre string `json:"nombre" binding:"required"`
		Descripcion string `json:"Descripcion" binding:"required"`
		Ingredientes string `json:"Ingredientes" binding:"required"`
		Fecha_creacion string `json:"Fecha_creacion" binding:"required"`
	}

	if err:= c.ShouldBindJSON(&data);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	perfumes_personalizados := entities.NewPerfumePersonalizado(data.Descripcion,data.Nombre,data.Ingredientes,data.Fecha_creacion)

	err := cpp_c.useCase.Execute(perfumes_personalizados)

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated,gin.H{"message":"Perfume personalizado creado"})
}