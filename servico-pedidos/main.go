package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"servico-pedidos/docs"
)

type Pedido struct {
	ID, UsuarioID, ProdutoID, Quantidade int
}

var db *sql.DB

func initDB() {
	var err error
	dbURL := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	db.Exec(`CREATE TABLE IF NOT EXISTS pedidos (id SERIAL PRIMARY KEY, usuario_id INT, produto_id INT, quantidade INT)`)
	log.Println("DB Pedidos OK")
}

func criarPedido(c *gin.Context) {
	var p Pedido
	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. Validar Usuário (Comunicação via HTTP)
	urlBaseUsuario := getEnv("URL_USUARIOS", "http://localhost:8081")
	respUsuario, err := http.Get(fmt.Sprintf("%s/usuarios/%d", urlBaseUsuario, req.UsuarioID))

	if err != nil || respUsuario.StatusCode != http.StatusOK {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado ou serviço offline"})
		return
	}
	respUsuario.Body.Close()

	// 2. Validar Produto (Comunicação via HTTP)
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
	err := db.QueryRow(`INSERT INTO pedidos (usuario_id, produto_id, quantidade) VALUES ($1, $2, $3) RETURNING id`, p.UsuarioID, p.ProdutoID, p.Quantidade).Scan(&p.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar"})
		return
	}
	c.JSON(http.StatusCreated, p)
}

func main() {
	initDB()
	defer db.Close()
	docs.SwaggerInfo.Host = ""
	router := gin.Default()
	router.POST("/pedidos", criarPedido)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8083")
}
