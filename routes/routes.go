package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jacolpn/go-api-gin/controllers"
	docs "github.com/jacolpn/go-api-gin/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()

	docs.SwaggerInfo_swagger.BasePath = "/"

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

	// Swagger.
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":5000")
}
