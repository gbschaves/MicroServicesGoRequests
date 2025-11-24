package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func proxyRequest(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro na configuração do Proxy"})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)

		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host

			// Remove o prefixo /api para mandar para o microserviço
			// Ex: Gateway recebe /api/usuarios -> Microserviço recebe /usuarios
			originalPath := c.Request.URL.Path
			req.URL.Path = strings.TrimPrefix(originalPath, "/api")
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func proxySwaggerRequest(target string, routePrefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro na configuração do Proxy"})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)

		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host

			originalPath := c.Request.URL.Path
			trimmedPath := strings.TrimPrefix(originalPath, routePrefix)
			req.URL.Path = "/swagger" + trimmedPath
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	targetUsuarios := getEnv("USER_SERVICE_URL", "http://localhost:8081")
	targetProdutos := getEnv("PRODUCT_SERVICE_URL", "http://localhost:8082")
	targetPedidos := getEnv("ORDER_SERVICE_URL", "http://localhost:8083")

	router := gin.Default()

	// --- CORS Middleware ---
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// === Rotas do Swagger ===
	// Estas rotas devem ser definidas antes ou de forma que não conflitem
	const docsUsuariosPrefix = "/api/usuarios/docs"
	router.GET(docsUsuariosPrefix+"/*proxyPath", proxySwaggerRequest(targetUsuarios, docsUsuariosPrefix))

	const docsProdutosPrefix = "/api/produtos/docs"
	router.GET(docsProdutosPrefix+"/*proxyPath", proxySwaggerRequest(targetProdutos, docsProdutosPrefix))

	const docsPedidosPrefix = "/api/pedidos/docs"
	router.GET(docsPedidosPrefix+"/*proxyPath", proxySwaggerRequest(targetPedidos, docsPedidosPrefix))

	// === Rotas da API (Específicas) ===
	// Removemos o router.Any com *path para evitar conflito com o Swagger

	// Usuários
	router.POST("/api/usuarios", proxyRequest(targetUsuarios))
	router.GET("/api/usuarios/:id", proxyRequest(targetUsuarios))

	// Produtos
	router.POST("/api/produtos", proxyRequest(targetProdutos))
	router.GET("/api/produtos/:id", proxyRequest(targetProdutos))

	// Pedidos
	router.POST("/api/pedidos", proxyRequest(targetPedidos))

	porta := getEnv("SERVER_PORT", ":8080")

	fmt.Printf("API Gateway rodando na porta %s\n", porta)
	router.Run(porta)
}
