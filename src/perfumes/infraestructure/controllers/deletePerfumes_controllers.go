package controllers



import (
	"ApiProduct/src/perfumes/application"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)

type DeletePerfumeController struct {
	useCase *application.DeletePerfume
}

func NewDeletePerfumeController(useCase *application.DeletePerfume) *DeletePerfumeController {
	return &DeletePerfumeController{useCase: useCase}
}

	func (ctrl *DeletePerfumeController) Execute(c *gin.Context){
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}
	
		err = ctrl.useCase.Execute(id) // Ahora pasamos el ID correctamente
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{"message": "Perfume eliminado con éxito"})
	}
	

