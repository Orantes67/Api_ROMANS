package infraestructure


import (

	"ApiProduct/src/perfumes/application"
	"ApiProduct/src/perfumes/infraestructure/controllers"
	"github.com/gin-gonic/gin"
 )

func Init(router *gin.Engine){

	ps := NewMySQL()

	createPerfumeService := application.NewPerfume(ps)
	listPerfumeService := application.NewListPerfume(ps)
	updatePerfumeService := application.NewUpdatePerfume(ps)
	deletePerfumeService := application.NewDeletePerfume(ps)


	cretePerfumeServiceController := controllers.NewCreatePefumesController(createPerfumeService)
	listPerfumeServiceController := controllers.NewListPerfumeController(listPerfumeService)
	updatePerfumeServiceController := controllers.NewUpdatePerfumeController(updatePerfumeService)
	deletePerfumeServiceController := controllers.NewDeletePerfumeController(deletePerfumeService)
	
	PerfumesRoutes(router, PerfumeHandler{
		Create: cretePerfumeServiceController, 
		Get:    listPerfumeServiceController,  
		Edit:   updatePerfumeServiceController,
		Delete: deletePerfumeServiceController,
	})

}