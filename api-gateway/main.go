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

// Função auxiliar para ler variáveis de ambiente com valor padrão
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

			// Mantém o path original, removendo apenas o prefixo /api se necessário
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
	// --- CORREÇÃO AQUI: ---
	// O Gateway não conecta no banco. Removemos a parte do sql.Open.
	// Ele lê as URLs dos outros serviços via variáveis de ambiente (definidas no docker-compose).

	// Se não encontrar a variável (ex: rodando local no PC), usa localhost como fallback.
	targetUsuarios := getEnv("USER_SERVICE_URL", "http://localhost:8081")
	targetProdutos := getEnv("PRODUCT_SERVICE_URL", "http://localhost:8082")
	targetPedidos := getEnv("ORDER_SERVICE_URL", "http://localhost:8083")

	router := gin.Default()

	// === Rotas do Swagger ===
	const docsUsuariosPrefix = "/api/usuarios/docs"
	router.GET(docsUsuariosPrefix+"/*proxyPath", proxySwaggerRequest(targetUsuarios, docsUsuariosPrefix))

	const docsProdutosPrefix = "/api/produtos/docs"
	router.GET(docsProdutosPrefix+"/*proxyPath", proxySwaggerRequest(targetProdutos, docsProdutosPrefix))

	const docsPedidosPrefix = "/api/pedidos/docs"
	router.GET(docsPedidosPrefix+"/*proxyPath", proxySwaggerRequest(targetPedidos, docsPedidosPrefix))

	// === Rotas da API ===
	// Note que usamos Any para passar POST, GET, PUT, DELETE, etc.
	router.Any("/api/usuarios/*path", proxyRequest(targetUsuarios))
	// O "*path" pega qualquer coisa depois, ex: /api/usuarios/1 ou /api/usuarios
	// Mas como você definiu rotas específicas antes, vamos manter o padrão simples:

	router.POST("/api/usuarios", proxyRequest(targetUsuarios))
	router.GET("/api/usuarios/:id", proxyRequest(targetUsuarios))

	router.POST("/api/produtos", proxyRequest(targetProdutos))
	router.GET("/api/produtos/:id", proxyRequest(targetProdutos))

	router.POST("/api/pedidos", proxyRequest(targetPedidos))

	// A porta padrão dentro do container será 8080
	porta := getEnv("SERVER_PORT", ":8080")

	fmt.Printf("API Gateway rodando na porta %s\n", porta)
	fmt.Printf("Conectando Usuários em: %s\n", targetUsuarios)
	fmt.Printf("Conectando Produtos em: %s\n", targetProdutos)
	fmt.Printf("Conectando Pedidos em: %s\n", targetPedidos)

	router.Run(porta)
}
