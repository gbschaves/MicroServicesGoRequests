package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// proxyRequest (Esta é a função que já tínhamos para a API)
func proxyRequest(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse(target)
		if err != nil {
			panic(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)

		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host

			originalPath := c.Request.URL.Path
			// Ex: /api/usuarios/1 -> /usuarios/1
			req.URL.Path = strings.TrimPrefix(originalPath, "/api")
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

// proxySwaggerRequest (Função de proxy para o Swagger)
// Ele substitui o prefixo da rota do gateway pelo prefixo /swagger do serviço
func proxySwaggerRequest(target string, routePrefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// O target é o microserviço (ex: http://localhost:8081)
		remote, err := url.Parse(target)
		if err != nil {
			panic(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)

		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host

			// Pega o caminho original (ex: /api/usuarios/docs/index.html)
			originalPath := c.Request.URL.Path

			// Remove o prefixo da rota do gateway (ex: /index.html)
			trimmedPath := strings.TrimPrefix(originalPath, routePrefix)

			// Adiciona o prefixo /swagger que o serviço espera
			req.URL.Path = "/swagger" + trimmedPath
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

// === FUNÇÃO MAIN ATUALIZADA ===
func main() {
	router := gin.Default()

	targetUsuarios := "http://localhost:8081"
	targetProdutos := "http://localhost:8082"
	targetPedidos := "http://localhost:8083"

	// === Rotas do Swagger ===
	const docsUsuariosPrefix = "/api/usuarios/docs"
	router.GET(docsUsuariosPrefix+"/*proxyPath", proxySwaggerRequest(targetUsuarios, docsUsuariosPrefix))

	const docsProdutosPrefix = "/api/produtos/docs"
	router.GET(docsProdutosPrefix+"/*proxyPath", proxySwaggerRequest(targetProdutos, docsProdutosPrefix))

	// NOVA ROTA:
	const docsPedidosPrefix = "/api/pedidos/docs"
	router.GET(docsPedidosPrefix+"/*proxyPath", proxySwaggerRequest(targetPedidos, docsPedidosPrefix))

	// Rotas da API
	router.POST("/api/usuarios", proxyRequest(targetUsuarios))
	router.GET("/api/usuarios/:id", proxyRequest(targetUsuarios))

	router.POST("/api/produtos", proxyRequest(targetProdutos))
	router.GET("/api/produtos/:id", proxyRequest(targetProdutos))

	router.POST("/api/pedidos", proxyRequest(targetPedidos))

	porta := ":8080"
	fmt.Printf("API Gateway (Atualizado) rodando na porta %s\n", porta)
	fmt.Printf("Acesse a Doc. de Usuários em: http://localhost%s%s/index.html\n", porta, docsUsuariosPrefix)
	fmt.Printf("Acesse a Doc. de Produtos em: http://localhost%s%s/index.html\n", porta, docsProdutosPrefix)
	// MENSAGEM ATUALIZADA:
	fmt.Printf("Acesse a Doc. de Pedidos em: http://localhost%s%s/index.html\n", porta, docsPedidosPrefix)
	router.Run(porta)
}
