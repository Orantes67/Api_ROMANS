package controllers


import (
	"ApiProduct/src/perfumes_personalizados/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListPerfumePersonalizadoController struct{
	useCase *application.ListPerfumePersonalizados
}

func NewListPerfumePersonalizadoController(useCase *application.ListPerfumePersonalizados)*ListPerfumePersonalizadoController{
	return &ListPerfumePersonalizadoController{useCase: useCase}
}


func (ctrl *ListPerfumePersonalizadoController) Execute(c *gin.Context){
	perfumespersonalizados, err := ctrl.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,perfumespersonalizados)
}