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
	"servico-usuarios/docs" // Importe a pasta docs
	"strconv"
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
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	db.Exec(`CREATE TABLE IF NOT EXISTS usuarios (id SERIAL PRIMARY KEY, nome TEXT, email TEXT)`)
	log.Println("DB Usuários OK")
}

func criarUsuario(c *gin.Context) {
	var u Usuario
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.QueryRow(`INSERT INTO usuarios (nome, email) VALUES ($1, $2) RETURNING id`, u.Nome, u.Email).Scan(&u.ID)
	c.JSON(http.StatusCreated, u)
}

func buscarUsuarioPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var u Usuario
	err := db.QueryRow(`SELECT id, nome, email FROM usuarios WHERE id = $1`, id).Scan(&u.ID, &u.Nome, &u.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Não encontrado"})
		return
	}
	c.JSON(http.StatusOK, u)
}

func main() {
	initDB()
	defer db.Close()

	// O PULO DO GATO: Deixar vazio permite que o Swagger use o host do navegador (seu dominio)
	docs.SwaggerInfo.Host = ""

	router := gin.Default()
	router.POST("/usuarios", criarUsuario)
	router.GET("/usuarios/:id", buscarUsuarioPorID)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8081")
}
