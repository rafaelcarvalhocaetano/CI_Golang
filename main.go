package main

import (
	"github.com/rafaelcarvalhocaetano/ci_cd_golang/database"
	"github.com/rafaelcarvalhocaetano/ci_cd_golang/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
