# 使用官方 golang 镜像作为构建环境
FROM --platform=$BUILDPLATFORM golang:1.20-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go mod 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
ARG TARGETARCH
RUN GOOS=linux GOARCH=$TARGETARCH CGO_ENABLED=0 go build -o main .

# 使用轻量级的 alpine 作为运行环境
FROM --platform=$TARGETPLATFORM alpine:latest

# 安装 ca-certificates
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建阶段复制编译好的应用
COPY --from=builder /app/main .
# 复制模板和静态文件
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

# 暴露端口
EXPOSE 8914

# 设置 gin 为生产模式
ENV GIN_MODE=release

# 运行应用
CMD ["./main"] 