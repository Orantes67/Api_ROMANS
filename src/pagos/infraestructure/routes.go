package infraestructure

import (
	"ApiProduct/src/pagos/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

type PagosHandler struct{
	Create *controllers.CreatePagosController
	Get *controllers.ListPagosController
	Edit *controllers.UpdatePagosController
	Delete *controllers.DeletePagosController
}

func PagosRoutes(router *gin.Engine,handler PagosHandler){
	pagosGroup := router.Group("/pagos")
	{
		pagosGroup.POST("/", handler.Create.Execute)
		pagosGroup.GET("/", handler.Get.Execute)
		pagosGroup.PUT("/:id", handler.Edit.Execute)
		pagosGroup.DELETE("/:id", handler.Delete.Execute)
	}
}