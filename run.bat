@echo off
echo Iniciando todos os servicos...

REM Define um titulo para cada janela para sabermos quem e quem

start "API Gateway (8080)" cmd /c "cd api-gateway && go run main.go"
start "Servico Usuarios (8081)" cmd /c "cd servico-usuarios && go run main.go"
start "Servico Produtos (8082)" cmd /c "cd servico-produtos && go run main.go"
start "Servico Pedidos (8083)" cmd /c "cd servico-pedidos && go run main.go"

echo.
echo Servicos iniciados em janelas separadas.
echo Para parar tudo, feche as 4 janelas que foram abertas.