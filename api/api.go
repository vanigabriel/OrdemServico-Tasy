package api

import (
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanigabriel/OrdemServico-Tasy/entity"
	ordem "github.com/vanigabriel/OrdemServico-Tasy/usecases"
)

func SetupRouter(service *ordem.Service) *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	r.POST("/ordemservico", PostOS(service))
	r.POST("/ordemservico/:os/files", PostFiles(service))

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

		ordem, err := service.InsertOS(&OS)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Processado sem erros.", "numero": ordem})
	}
}

func PostFiles(service *ordem.Service) func(c *gin.Context) {
	return func(c *gin.Context) {

		var ordem string
		ordem = c.Param("os")
		if len(ordem) == 0 {
			log.Println("Número da OS não informada")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parametro OS não informado"})
			return
		}

		multipart, err := c.Request.MultipartReader()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for {
			mimePart, err := multipart.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			_, params, err := mime.ParseMediaType(mimePart.Header.Get("Content-Disposition"))
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			//Get File
			file, err := ioutil.ReadAll(mimePart)
			if err != nil {
				log.Fatal(err)
			}

			err = service.InsertAnexos(ordem, params["filename"], file)
			if err != nil {
				log.Fatal(err)
			}

		}

		c.JSON(http.StatusCreated, gin.H{"message": "Processado sem erros."})
	}
}
