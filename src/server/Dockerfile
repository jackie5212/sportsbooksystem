FROM golang:1.21-alpine AS builder

WORKDIR /app

# 安装依赖
RUN apk add --no-cache git

# 设置 Go 模块代理为国内镜像
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=sum.golang.org

# 复制go.mod和go.sum
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 编译
RUN CGO_ENABLED=0 GOOS=linux go build -o tennis-booking cmd/main.go

# 最终镜像
FROM alpine:latest

WORKDIR /app

# 安装必要的工具
RUN apk add --no-cache wget ca-certificates tzdata

# 复制编译好的二进制文件
COPY --from=builder /app/tennis-booking .
COPY --from=builder /app/config.yaml.example ./config.yaml

# 创建日志目录
RUN mkdir -p /app/logs

# 暴露端口
EXPOSE 8080

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s --start-period=40s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# 运行
CMD ["./tennis-booking"]
