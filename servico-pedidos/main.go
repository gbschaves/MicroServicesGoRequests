package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"servico-pedidos/docs" // Importa a documentação

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// --- Estruturas ---
type Pedido struct {
	ID         int `json:"id"`
	UsuarioID  int `json:"usuario_id"`
	ProdutoID  int `json:"produto_id"`
	Quantidade int `json:"quantidade"`
}

type NovoPedidoRequest struct {
	UsuarioID  int `json:"usuario_id"`
	ProdutoID  int `json:"produto_id"`
	Quantidade int `json:"quantidade"`
}

var db *sql.DB

// --- Função Auxiliar (Que estava faltando) ---
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func initDB() {
	var err error
	dbURL := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Banco não responde:", err)
	}

	query := `CREATE TABLE IF NOT EXISTS pedidos (
		id SERIAL PRIMARY KEY,
		usuario_id INT NOT NULL,
		produto_id INT NOT NULL,
		quantidade INT NOT NULL
	)`
	if _, err = db.Exec(query); err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}
	log.Println("Banco de Pedidos OK!")
}

// @Summary      Cria um novo pedido
// @Router       /pedidos [post]
func criarPedido(c *gin.Context) {
	var req NovoPedidoRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. Validar Usuário (Comunicação via HTTP com o microsserviço de usuários)
	urlBaseUsuario := getEnv("URL_USUARIOS", "http://localhost:8081")
	respUsuario, err := http.Get(fmt.Sprintf("%s/usuarios/%d", urlBaseUsuario, req.UsuarioID))

	if err != nil || respUsuario.StatusCode != http.StatusOK {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado ou serviço offline"})
		return
	}
	respUsuario.Body.Close()

	// 2. Validar Produto (Comunicação via HTTP com o microsserviço de produtos)
	urlBaseProduto := getEnv("URL_PRODUTOS", "http://localhost:8082")
	respProduto, err := http.Get(fmt.Sprintf("%s/produtos/%d", urlBaseProduto, req.ProdutoID))

	if err != nil || respProduto.StatusCode != http.StatusOK {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado ou serviço offline"})
		return
	}
	respProduto.Body.Close()

	// 3. Salvar Pedido no Banco
	pedido := Pedido{
		UsuarioID:  req.UsuarioID,
		ProdutoID:  req.ProdutoID,
		Quantidade: req.Quantidade,
	}

	query := `INSERT INTO pedidos (usuario_id, produto_id, quantidade) VALUES ($1, $2, $3) RETURNING id`
	err = db.QueryRow(query, pedido.UsuarioID, pedido.ProdutoID, pedido.Quantidade).Scan(&pedido.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar pedido no banco"})
		return
	}

	c.JSON(http.StatusCreated, pedido)
}

func main() {
	initDB()
	defer db.Close()

	// CORREÇÃO PARA O SWAGGER (Permite funcionar via HTTPS/Cloudflare)
	docs.SwaggerInfo.Host = ""

	router := gin.Default()
	router.POST("/pedidos", criarPedido)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Porta 8083
	router.Run(":8083")
}
