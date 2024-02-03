FROM alpine:latest
LABEL MAINTAINER="IOMServer"

WORKDIR /go/src/server
COPY ./main ./
COPY ./res ./res
COPY ./index.html ./

RUN chmod 755 ./main

EXPOSE 10000
EXPOSE 8088

ENTRYPOINT ./main