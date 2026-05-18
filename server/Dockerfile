FROM golang:1.21-alpine AS builder

WORKDIR /app

# 安装依赖
RUN apk add --no-cache git

# 复制go.mod和go.sum
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 编译
RUN CGO_ENABLED=0 GOOS=linux go build -o tennis-booking cmd/main.go

# 最终镜像
FROM alpine:latest

WORKDIR /root/

# 复制编译好的二进制文件
COPY --from=builder /app/tennis-booking .
COPY --from=builder /app/config.yaml.example ./config.yaml

# 暴露端口
EXPOSE 8080

# 运行
CMD ["./tennis-booking"]
