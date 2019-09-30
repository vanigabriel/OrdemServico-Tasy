package api

import (
	"io"
	"io/ioutil"
	"mime"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanigabriel/OrdemServico-Tasy/entity"
	ordem "github.com/vanigabriel/OrdemServico-Tasy/usecases"

	"github.com/evalphobia/go-timber/timber"
)

// SetupRouter cria roteamento
func SetupRouter(service *ordem.Service, log *timber.Client) *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	r.POST("/ordemservico", PostOS(service, log))
	r.POST("/ordemservico/:os/files", PostFiles(service, log))

	return r
}

// PostOS cria ordem de serviço
func PostOS(service *ordem.Service, log *timber.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var OS entity.OrdemServico

		err := c.BindJSON(&OS)
		if err != nil {
			log.Err(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ordem, err := service.InsertOS(&OS)
		if err != nil {
			log.Err(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Processado sem erros.", "numero": ordem})
	}
}

// PostFiles adiciona anexos à ordem de serviço
func PostFiles(service *ordem.Service, log *timber.Client) func(c *gin.Context) {
	return func(c *gin.Context) {

		var ordem string
		ordem = c.Param("os")
		if len(ordem) == 0 {
			log.Err("Número da OS não informada")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parametro OS não informado"})
			return
		}

		multipart, err := c.Request.MultipartReader()
		if err != nil {
			log.Err(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for {
			mimePart, err := multipart.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Err(err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			_, params, err := mime.ParseMediaType(mimePart.Header.Get("Content-Disposition"))
			if err != nil {
				log.Err(err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			//Get File
			file, err := ioutil.ReadAll(mimePart)
			if err != nil {
				log.Err(err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			err = service.InsertAnexos(ordem, params["filename"], file)
			if err != nil {
				log.Err(err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

		}

		c.JSON(http.StatusCreated, gin.H{"message": "Processado sem erros."})
	}
}
