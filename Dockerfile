# 打包阶段 golang 版本规定
FROM golang:1.15 as builder

# 启用go mod
ENV GOMODULE111=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /web

COPY . .

# CGO_ENABLED禁用cgo 然后指定OS等，并go build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo .

FROM scratch

WORKDIR /web

COPY --from=builder /web  .

ENV GIN_MODE=release\
    PORT=8087

# 端口暴露
EXPOSE 8087
ENTRYPOINT ["./sso"]

