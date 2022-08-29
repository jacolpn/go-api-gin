package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jacolpn/go-api-gin/controllers"
	"github.com/jacolpn/go-api-gin/database"
	"github.com/jacolpn/go-api-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	// Muda as descrições que aparecem no console.
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()

	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno teste", CPF: "00000000000", RG: "000000000"}

	database.DB.Create(&aluno)

	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno

	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	// t.Fatalf("Teste falhou de propósito, não se preocupe!")
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)

	req, _ := http.NewRequest("GET", "/jackson", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	// Validação sem o testify assert.
	// if resposta.Code != http.StatusOK {
	// 	t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	// }

	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")

	mockDaResposta := `{"API diz:":"E ai jackson, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestListaTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()

	// Após fazer tudo, vai deletar o aluno.
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
	// fmt.Println(resposta.Body)
}

func TestBuscaAlunoPorCPF(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()

	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/00000000000", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorIdHandle(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()

	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)

	// fmt.Println(alunoMock.Nome)

	assert.Equal(t, "Aluno teste", alunoMock.Nome, "Os nome's devem ser iguais")
	assert.Equal(t, "00000000000", alunoMock.CPF, "Os CPF's devem ser iguais")
	assert.Equal(t, "000000000", alunoMock.RG, "Os RG's devem ser iguais")
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDaBusca, nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()

	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)

	aluno := models.Aluno{Nome: "Aluno teste", CPF: "47000000000", RG: "700000000"}
	valorJson, _ := json.Marshal(aluno)
	pathParaEditar := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)

	assert.Equal(t, "47000000000", alunoMockAtualizado.CPF)
	assert.Equal(t, "700000000", alunoMockAtualizado.RG)
	assert.Equal(t, "Aluno teste", alunoMockAtualizado.Nome)
	// fmt.Print(resposta.Body)
}
