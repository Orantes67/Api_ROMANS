package infraestructure

import (
	"ApiProduct/src/perfumes_personalizados/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

type PerfumePersonalizadoHandler struct{
	Create *controllers.CreatePerfumePersonalizadoController
	Get *controllers.ListPerfumePersonalizadoController
	Edit *controllers.UpdatePefumePersonalizadoController
	Delete *controllers.DeletePerfumesPersonalizadoController
}

func PerfumesPersonalizadosRoutes(router *gin.Engine, handler PerfumePersonalizadoHandler){
	perfumesPersonalizadosGroup := router.Group("/perfumes")
	{
		perfumesPersonalizadosGroup.POST("/",handler.Create.Execute)
		perfumesPersonalizadosGroup.GET("/",handler.Get.Execute)
		perfumesPersonalizadosGroup.PUT("/:id",handler.Edit.Execute)
		perfumesPersonalizadosGroup.DELETE("/:id",handler.Delete.Execute)
	}
}