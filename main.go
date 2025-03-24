package main


import (
	"ApiProduct/src/pagos/infraestructure"
    perfumesInfra "ApiProduct/src/perfumes/infraestructure"
	"github.com/gin-gonic/gin"
	perfumespersonalizadosInfra"ApiProduct/src/perfumes_personalizados/infraestructure"
)
	

func main() {
	r := gin.Default()
	infraestructure.Init(r)
	perfumesInfra.Init(r)
	perfumespersonalizadosInfra.Init(r)
	r.Run(":8080")
	
}