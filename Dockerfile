# 使用官方 golang 镜像作为构建环境
FROM golang:1.20-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装 gcc 和必要的工具
RUN apk add --no-cache gcc musl-dev

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 编译
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

# 使用更小的基础镜像
FROM alpine:latest

# 安装 sqlite 和 CA 证书
RUN apk add --no-cache sqlite ca-certificates

# 设置工作目录
WORKDIR /app

# 创建数据目录
RUN mkdir -p /app/data

# 从 builder 阶段复制编译好的程序
COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

# 设置数据目录的权限
RUN chown -R nobody:nobody /app/data

# 使用非 root 用户运行
USER nobody

# 暴露端口
EXPOSE 8914

# 声明数据卷
VOLUME ["/app/data"]

# 运行程序
CMD ["./main"] 