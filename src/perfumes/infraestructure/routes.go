package infraestructure

import (
	"ApiProduct/src/perfumes/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

type PerfumeHandler struct {
	Create   *controllers.CreatePerfumesController
	Get *controllers.ListPerfumeController
	Edit *controllers.UpdatePerfumeController
	Delete *controllers.DeletePerfumeController
}

func PerfumesRoutes(router *gin.Engine, handler PerfumeHandler) {
	perfumesGroup := router.Group("/perfumes")
	{
		perfumesGroup.POST("/", handler.Create.Execute)
		perfumesGroup.GET("/", handler.Get.Execute)
		perfumesGroup.PUT("/:id", handler.Edit.Execute)
		perfumesGroup.DELETE("/:id", handler.Delete.Execute)
	}
}