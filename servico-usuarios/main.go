package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "servico-usuarios/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

var db *sql.DB

func initDB() {
	var err error
	dbURL := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Banco não responde:", err)
	}

	// Cria a tabela se não existir
	query := `CREATE TABLE IF NOT EXISTS usuarios (
		id SERIAL PRIMARY KEY,
		nome TEXT NOT NULL,
		email TEXT NOT NULL
	)`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}
	log.Println("Banco de Dados de Usuários conectado e configurado!")
}

// @Summary      Cria um novo usuário
// @Router       /usuarios [post]
func criarUsuario(c *gin.Context) {
	var u Usuario
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// INSERE NO BANCO DE DADOS REAL
	query := `INSERT INTO usuarios (nome, email) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(query, u.Nome, u.Email).Scan(&u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar no banco"})
		return
	}

	c.JSON(http.StatusCreated, u)
}

// @Summary      Busca um usuário por ID
// @Router       /usuarios/{id} [get]
func buscarUsuarioPorID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var u Usuario
	// BUSCA DO BANCO DE DADOS REAL
	query := `SELECT id, nome, email FROM usuarios WHERE id = $1`
	row := db.QueryRow(query, id)

	err := row.Scan(&u.ID, &u.Nome, &u.Email)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

func main() {
	initDB()
	defer db.Close()

	router := gin.Default()
	router.POST("/usuarios", criarUsuario)
	router.GET("/usuarios/:id", buscarUsuarioPorID)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8081")
}
