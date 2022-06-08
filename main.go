package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	application := app.Gerar()
	erro := application.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
