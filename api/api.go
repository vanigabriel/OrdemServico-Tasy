package api

import (
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vanigabriel/OrdemServico-Tasy/entity"
	ordem "github.com/vanigabriel/OrdemServico-Tasy/usecases"
)

func SetupRouter(service *ordem.Service) *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	r.POST("/ordemservico", PostOS(service))
	r.POST("/ordemservico/files", PostFiles)

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

func PostFiles(c *gin.Context) {
	multipart, err := c.Request.MultipartReader()
	if err != nil {
		log.Fatalln("Failed to create MultipartReader", err)
	}

	for {
		mimePart, err := multipart.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading multipart section: %v", err)
			break
		}
		_, params, err := mime.ParseMediaType(mimePart.Header.Get("Content-Disposition"))
		if err != nil {
			log.Printf("Invalid Content-Disposition: %v", err)
			break
		}

		//Create File
		f, err := os.OpenFile(params["filename"], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		slurp, err := ioutil.ReadAll(mimePart)
		if err != nil {
			log.Fatal(err)
		}

		_, err = f.Write(slurp)
		if err != nil {
			log.Fatal(err)
		}

	}

	c.JSON(http.StatusCreated, gin.H{"message": "Processado sem erros."})
}
