FROM nginx:1.25-alpine
LABEL MAINTAINER="Dashboard"

COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY ./dist /usr/share/nginx/html/

#暴露容器8087端口
EXPOSE 8087
