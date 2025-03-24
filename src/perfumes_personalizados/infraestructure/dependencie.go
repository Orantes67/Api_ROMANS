package infraestructure

import (
	"ApiProduct/src/perfumes_personalizados/application"
	"ApiProduct/src/perfumes_personalizados/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine){

	ps := NewMySQL()

	createPerfumePersonalizadoService := application.NewPerfumePersonalizado(ps)
	listPerfumePersonalizadoService := application.NewListPerfumePersonalizado(ps)
	updatePerfumePersonalizadoService := application.NewUpdatePerfumePersonalizado(ps)
	deletePerfumePersonalizadoService := application.NewDeletePerfumePersonalizado(ps)

	createPerfumePersonalizadoServiceController := controllers.NewCreatePefumesPersonalizadoController(createPerfumePersonalizadoService)
	listPerfumePersonalizadoServiceController := controllers.NewListPerfumePersonalizadoController(listPerfumePersonalizadoService)
	updatePerfumePersonalizadoServiceController := controllers.NewUpdatePerfumePersonalizadoController(updatePerfumePersonalizadoService)
	deletePerfumePersonalizadoServiceController := controllers.NewDeletePerfumesPersonalizadoController(deletePerfumePersonalizadoService)

	PerfumesPersonalizadosRoutes(router, PerfumePersonalizadoHandler{
		Create: createPerfumePersonalizadoServiceController,
		Get: listPerfumePersonalizadoServiceController,
		Edit: updatePerfumePersonalizadoServiceController,
		Delete: deletePerfumePersonalizadoServiceController,
	})
}