package main

import (
	"fmt"
	"net/http"
	"strconv"

	_ "servico-produtos/docs" // <-- IMPORTANTE

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 1. Nosso Modelo de Dados
type Produto struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

// 2. "Banco de Dados"
var produtos []Produto
var proximoID = 1

// 3. Handlers (com anotações Swagger)

// @Summary      Cria um novo produto
// @Description  Cria um produto com nome e preço
// @Tags         produtos
// @Accept       json
// @Produce      json
// @Param        produto  body      Produto  true  "Informações do Produto"
// @Success      201      {object}  Produto
// @Failure      400      {object}  map[string]string
// @Router       /produtos [post]
func criarProduto(c *gin.Context) {
	var novoProduto Produto
	if err := c.BindJSON(&novoProduto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	novoProduto.ID = proximoID
	proximoID++
	produtos = append(produtos, novoProduto)
	c.JSON(http.StatusCreated, novoProduto)
}

// @Summary      Busca um produto por ID
// @Description  Retorna os dados de um produto específico
// @Tags         produtos
// @Produce      json
// @Param        id   path      int  true  "ID do Produto"
// @Success      200  {object}  Produto
// @Failure      400      {object}  map[string]string
// @Failure      404      {object}  map[string]string
// @Router       /produtos/{id} [get]
func buscarProdutoPorID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	for _, p := range produtos {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
}

// 4. Função Principal
// @title        API de Produtos (Microserviço)
// @version      1.0
// @description  Este é o microserviço de produtos.
// @host         localhost:8080
// @BasePath     /api/produtos
func main() {
	router := gin.Default()

	router.POST("/produtos", criarProduto)
	router.GET("/produtos/:id", buscarProdutoPorID)

	// Rota do Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	porta := ":8082"
	fmt.Printf("Serviço de Produtos (com Swagger) rodando na porta %s\n", porta)
	fmt.Printf("Acesse a documentação em http://localhost:8082/swagger/index.html\n")
	router.Run(porta)
}
