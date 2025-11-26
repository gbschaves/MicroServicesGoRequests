FROM nginx:alpine

# Remove a configuração padrão do Nginx
RUN rm /etc/nginx/conf.d/default.conf

# Copia a nossa configuração personalizada
COPY nginx.conf /etc/nginx/conf.d/

# Copia os arquivos do site
RUN rm -rf /usr/share/nginx/html/*
COPY *.html /usr/share/nginx/html/
# Se tiver imagens, descomente:
# COPY *.png /usr/share/nginx/html/

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]