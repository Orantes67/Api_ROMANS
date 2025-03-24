package controllers


import (
	"ApiProduct/src/perfumes_personalizados/application"
	"ApiProduct/src/perfumes_personalizados/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)

type UpdatePefumePersonalizadoController struct{
	useCase *application.UpdatePerfumePersonalizados
}

func NewUpdatePerfumePersonalizadoController(useCase *application.UpdatePerfumePersonalizados) *UpdatePefumePersonalizadoController{
	return &UpdatePefumePersonalizadoController{useCase: useCase}

}

func (ctrl *UpdatePefumePersonalizadoController) Execute(c *gin.Context){
	idParam := c.Param("id")
	id,err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "ID invalido"})
		return
	}

	var data struct{
		Nombre string `json:"nombre" binding:"required"`
		Descripcion string `json:"Descripcion" binding:"required"`
		Ingredientes string `json:"Ingredientes" binding:"required"`
		Fecha_creacion string `json:"Fecha_creacion" binding:"required"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatePerfumePersonalizado := entities.NewPerfumePersonalizado(data.Descripcion,data.Ingredientes,data.Nombre,data.Fecha_creacion)
	err = ctrl.useCase.Execute(id,updatePerfumePersonalizado)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Perfume personalizado actualizado con éxito"})
}
