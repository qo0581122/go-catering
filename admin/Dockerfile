FROM golang:alpine as builder

WORKDIR /root/admin
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="jiajie.fu"

WORKDIR /root/admin

COPY --from=0 /root/admin/server ./
COPY --from=0 /root/admin/conf/config.yaml ./conf/

EXPOSE 8888
ENTRYPOINT ./server
