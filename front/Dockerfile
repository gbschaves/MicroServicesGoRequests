FROM nginx:alpine
# Remove a config padrão do Nginx
RUN rm -rf /usr/share/nginx/html/*
# Copia seus arquivos HTML para o servidor
COPY *.html /usr/share/nginx/html/
COPY *.png /usr/share/nginx/html/
# Configuração customizada para permitir rotas limpas (opcional, mas recomendado)
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]