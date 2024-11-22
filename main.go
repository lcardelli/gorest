package main

import (
	"fmt"

	"github.com/lcardelli/gorest/database"
	"github.com/lcardelli/gorest/models"
	"github.com/lcardelli/gorest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "João", Historia: "Historia do João"},
		{Id: 2, Nome: "Maria", Historia: "Historia da Maria"},
	}
	database.Connect()
	fmt.Println("Servidor iniciado na porta 8080!")
	routes.HandleRequests()
}
