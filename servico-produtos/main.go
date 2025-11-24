package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "servico-produtos/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Produto struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
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

	query := `CREATE TABLE IF NOT EXISTS produtos (
		id SERIAL PRIMARY KEY,
		nome TEXT NOT NULL,
		preco NUMERIC(10, 2) NOT NULL
	)`
	if _, err = db.Exec(query); err != nil {
		log.Fatal(err)
	}
	log.Println("Banco de Produtos OK!")
}

// @Summary      Cria um novo produto
// @Router       /produtos [post]
func criarProduto(c *gin.Context) {
	var p Produto
	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO produtos (nome, preco) VALUES ($1, $2) RETURNING id`
	if err := db.QueryRow(query, p.Nome, p.Preco).Scan(&p.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, p)
}

// @Summary      Busca um produto por ID
// @Router       /produtos/{id} [get]
func buscarProdutoPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var p Produto

	query := `SELECT id, nome, preco FROM produtos WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&p.ID, &p.Nome, &p.Preco)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func main() {
	initDB()
	defer db.Close()

	router := gin.Default()
	router.POST("/produtos", criarProduto)
	router.GET("/produtos/:id", buscarProdutoPorID)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8082")
}
