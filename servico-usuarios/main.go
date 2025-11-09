package main

import (
	"fmt"
	"net/http"
	"strconv"

	_ "servico-usuarios/docs" // <-- IMPORTANTE (para o doc.json)

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 1. Nosso Modelo de Dados
type Usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// 2. Nosso "Banco de Dados"
var usuarios []Usuario
var proximoID = 1

// 3. Handlers (com anotações Swagger)

// @Summary      Cria um novo usuário
// @Description  Cria um usuário com nome e email
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Param        usuario  body      Usuario  true  "Informações do Usuário"
// @Success      201      {object}  Usuario
// @Failure      400      {object}  map[string]string
// @Router       /usuarios [post]
func criarUsuario(c *gin.Context) {
	var novoUsuario Usuario
	if err := c.BindJSON(&novoUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	novoUsuario.ID = proximoID
	proximoID++
	usuarios = append(usuarios, novoUsuario)
	c.JSON(http.StatusCreated, novoUsuario)
}

// @Summary      Busca um usuário por ID
// @Description  Retorna os dados de um usuário específico
// @Tags         usuarios
// @Produce      json
// @Param        id   path      int  true  "ID do Usuário"
// @Success      200  {object}  Usuario
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /usuarios/{id} [get]
func buscarUsuarioPorID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	for _, u := range usuarios {
		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
}

// 4. Função Principal
// @title        API de Usuários (Microserviço)
// @version      1.0
// @description  Este é o microserviço de usuários.
// @host         localhost:8080
// @BasePath     /api/usuarios
func main() {
	router := gin.Default()

	router.POST("/usuarios", criarUsuario)
	router.GET("/usuarios/:id", buscarUsuarioPorID)

	// Rota do Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	porta := ":8081"
	fmt.Printf("Serviço de Usuários (com Swagger) rodando na porta %s\n", porta)
	fmt.Printf("Acesse a documentação em http://localhost:8081/swagger/index.html\n")
	router.Run(porta)
}
