package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "servico-pedidos/docs" // <-- IMPORTANTE

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 1. Nossos Modelos de Dados
type Pedido struct {
	ID         int `json:"id"`
	UsuarioID  int `json:"usuario_id"`
	ProdutoID  int `json:"produto_id"`
	Quantidade int `json:"quantidade"`
}

// Struct apenas para o body do Swagger
type NovoPedidoRequest struct {
	UsuarioID  int `json:"usuario_id"`
	ProdutoID  int `json:"produto_id"`
	Quantidade int `json:"quantidade"`
}

// Structs "dummy" apenas para checar se a resposta dos outros serviços foi OK
type Usuario struct {
	ID int `json:"id"`
}
type Produto struct {
	ID int `json:"id"`
}

// 2. "Banco de Dados"
var pedidos []Pedido
var proximoID = 1
var gatewayURL = "http://localhost:8080"

// 3. Handler (com anotações Swagger)

// @Summary      Cria um novo pedido
// @Description  Cria um novo pedido validando usuário e produto
// @Tags         pedidos
// @Accept       json
// @Produce      json
// @Param        pedido  body      NovoPedidoRequest  true  "Informações do Pedido"
// @Success      201     {object}  Pedido
// @Failure      400     {object}  map[string]string
// @Failure      404     {object}  map[string]string
// @Router       /pedidos [post]
func criarPedido(c *gin.Context) {
	var req NovoPedidoRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 4. VERIFICA O USUÁRIO
	urlUsuario := fmt.Sprintf("%s/api/usuarios/%d", gatewayURL, req.UsuarioID)
	respUsuario, err := http.Get(urlUsuario)
	if err != nil || respUsuario.StatusCode != http.StatusOK {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}
	defer respUsuario.Body.Close()
	var usuario Usuario
	json.NewDecoder(respUsuario.Body).Decode(&usuario)

	// 5. VERIFICA O PRODUTO
	urlProduto := fmt.Sprintf("%s/api/produtos/%d", gatewayURL, req.ProdutoID)
	respProduto, err := http.Get(urlProduto)
	if err != nil || respProduto.StatusCode != http.StatusOK {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}
	defer respProduto.Body.Close()

	// 6. Se tudo deu certo, "salva" o pedido
	novoPedido := Pedido{
		ID:         proximoID,
		UsuarioID:  req.UsuarioID,
		ProdutoID:  req.ProdutoID,
		Quantidade: req.Quantidade,
	}
	proximoID++
	pedidos = append(pedidos, novoPedido)

	c.JSON(http.StatusCreated, novoPedido)
}

// 4. Função Principal
// @title        API de Pedidos (Microserviço)
// @version      1.0
// @description  Este é o microserviço de pedidos.
// @host         localhost:8080
// @BasePath     /api/pedidos
func main() {
	router := gin.Default()

	router.POST("/pedidos", criarPedido)

	// Rota do Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	porta := ":8083"
	fmt.Printf("Serviço de Pedidos (com Swagger) rodando na porta %s\n", porta)
	fmt.Printf("Acesse a documentação em http://localhost:8083/swagger/index.html\n")
	router.Run(porta)
}
