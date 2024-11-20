FROM nginx:1.27-alpine
COPY ./static /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY ./cert/ /etc/nginx/ssl/
