package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanigabriel/OrdemServico-Tasy/entity"
	ordem "github.com/vanigabriel/OrdemServico-Tasy/usecases"
)

func SetupRouter(service *ordem.Service) *gin.Engine {
	r := gin.Default()

	r.POST("/ordemservico", PostOS(service))

	return r
}

func PostOS(service *ordem.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		log.Println("Ininiciando rota")

		var OS entity.OrdemServico

		err := c.BindJSON(&OS)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = service.InsertOS(&OS)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Processado sem erros."})
	}
}
