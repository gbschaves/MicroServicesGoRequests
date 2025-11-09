#!/bin/bash

# Define o nome da nossa sessão do tmux
SESSION="microservicos"

echo "Iniciando sessão '$SESSION' no tmux..."

# 1. Cria a sessão com o primeiro serviço (Gateway)
#    -d = desconectado (em background)
#    -s = nome da sessão
#    -n = nome da "janela"
tmux new-session -d -s $SESSION -n "Gateway" 'cd api-gateway && go run main.go'

# 2. Adiciona os outros serviços em novas janelas
tmux new-window -t $SESSION -n "Usuarios" 'cd servico-usuarios && go run main.go'
tmux new-window -t $SESSION -n "Produtos" 'cd servico-produtos && go run main.go'
tmux new-window -t $SESSION -n "Pedidos" 'cd servico-pedidos && go run main.go'

echo "Serviços iniciados."
echo "Para ver os logs, conecte-se com: tmux attach -t $SESSION"
echo "Para parar tudo, rode: tmux kill-session -t $SESSION"