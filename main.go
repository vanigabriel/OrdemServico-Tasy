package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/vanigabriel/OrdemServico-Tasy/api"
	ordem "github.com/vanigabriel/OrdemServico-Tasy/usecases"
)

//PORT port to be used
const PORT = "9191"

var f, _ = os.OpenFile("log_"+time.Now().Format("01-02-2006")+".log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)

func main() {
	// Load env variables
	handleError(godotenv.Load())

	// Logging to a file.
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	defer f.Close()

	s := ordem.NewService(ordem.NewRepository())

	r := api.SetupRouter(s)

	r.Run(":" + PORT)

}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
