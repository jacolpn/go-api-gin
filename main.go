package main

import (
	"github.com/jacolpn/go-api-gin/database"
	"github.com/jacolpn/go-api-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	// models.Alunos = []models.Aluno{
	// 	{Nome: "Jackson Neves", CPF: "000.000.000-00", RG: "00000000000"},
	// 	{Nome: "Angélica Neves", CPF: "000.000.000-00", RG: "00000000000"},
	// }

	routes.HandleRequests()
}
