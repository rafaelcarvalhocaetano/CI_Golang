package main

import (
	// "log"
	// "github.com/joho/godotenv"
	"github.com/rafaelcarvalhocaetano/ci_cd_golang/database"
	"github.com/rafaelcarvalhocaetano/ci_cd_golang/routes"
)

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Some error occured. Err: %s", err)
	// }
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
