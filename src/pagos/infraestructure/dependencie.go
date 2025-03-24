package infraestructure

import (

	"ApiProduct/src/pagos/application"
	"ApiProduct/src/pagos/infraestructure/controllers"
	"github.com/gin-gonic/gin"
 )

 func Init(router *gin.Engine){

	ps := NewMySQL()

	createPagosService := application.NewPago(ps)
	listPagosService := application.NewListPagos(ps)
	updatePagosService := application.NewUpdatePago(ps)
	deletePagosService := application.NewDeletePagos(ps)


	cretePagosServiceController := controllers.NewCreatePagosController(createPagosService)
	listPagosServiceController := controllers.NewListPagosController(listPagosService)
	updatePagosServiceController := controllers.NewUpdatePagosController(updatePagosService)
	deletePagosServiceController := controllers.NewDeletePagosController(deletePagosService)
	
	PagosRoutes(router, PagosHandler{
		Create: cretePagosServiceController, 
		Get:    listPagosServiceController,  
		Edit:   updatePagosServiceController,
		Delete: deletePagosServiceController,
	})

}