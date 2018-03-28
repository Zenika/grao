FROM node:latest as builder

ADD . /usr/src/github.com/Zenika/rao/rao-front
WORKDIR /usr/src/github.com/Zenika/rao/rao-front

RUN npm install && npm rebuild node-sass && npm run build

FROM nginx:stable

WORKDIR /usr/src/github.com/Zenika/rao/rao-front
COPY --from=builder /usr/src/github.com/Zenika/rao/rao-front/dist/ /var/www/html/

COPY default.conf /etc/nginx/conf.d/
