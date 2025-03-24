package controllers


import (
	"ApiProduct/src/perfumes/application"
	"ApiProduct/src/perfumes/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePerfumesController struct {
	useCase *application.CreatePerfume
}

func NewCreatePefumesController(useCase *application.CreatePerfume) *CreatePerfumesController {
	return &CreatePerfumesController{useCase: useCase}
}

func (cp_c *CreatePerfumesController) Execute(c *gin.Context) {
	var data struct {
		Nombre    string `json:"nombre" binding:"required"`
		Descripcion string `json:"Descripcion" binding:"required"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	perfumes := entities.NewPerfume(data.Nombre, data.Descripcion )

	err := cp_c.useCase.Execute(perfumes) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Perfume creado con éxito"})
}
