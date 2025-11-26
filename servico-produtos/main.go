package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"servico-produtos/docs"
	"strconv"
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
	db.Exec(`CREATE TABLE IF NOT EXISTS produtos (id SERIAL PRIMARY KEY, nome TEXT, preco NUMERIC(10,2))`)
	log.Println("DB Produtos OK")
}

func criarProduto(c *gin.Context) {
	var p Produto
	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.QueryRow(`INSERT INTO produtos (nome, preco) VALUES ($1, $2) RETURNING id`, p.Nome, p.Preco).Scan(&p.ID)
	c.JSON(http.StatusCreated, p)
}

func buscarProdutoPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var p Produto
	err := db.QueryRow(`SELECT id, nome, preco FROM produtos WHERE id = $1`, id).Scan(&p.ID, &p.Nome, &p.Preco)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Não encontrado"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func main() {
	initDB()
	defer db.Close()
	docs.SwaggerInfo.Host = ""
	router := gin.Default()
	router.POST("/produtos", criarProduto)
	router.GET("/produtos/:id", buscarProdutoPorID)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8082")
}
