package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jacolpn/go-api-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()

	// Renderizar as páginas HTML.
	r.LoadHTMLGlob("templates/*")

	// Assets - CSS.
	r.Static("/assets", "./assets")

	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	r.POST("/alunos", controllers.CriaNovoAluno)

	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	r.PATCH("/alunos/:id", controllers.EditaAluno)

	// HTML.
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.RotaNaoEncontrada)

	r.Run(":5000")
}
