# 构建阶段
FROM node:18-alpine AS builder

WORKDIR /app

# 复制 package.json 和 package-lock.json
COPY package*.json ./

# 安装所有依赖（包括开发依赖）
RUN npm install

# 复制源代码
COPY . .

# 修复 node_modules 权限
RUN chmod -R +x node_modules/.bin

# 构建生产版本
RUN npm run build

# 生产阶段 - 使用 Nginx
FROM nginx:alpine

# 复制自定义 nginx 配置
COPY nginx.conf /etc/nginx/conf.d/default.conf

# 复制构建的文件到 nginx 目录
COPY --from=builder /app/dist /usr/share/nginx/html

# 暴露端口
EXPOSE 80

# 启动 Nginx
CMD ["nginx", "-g", "daemon off;"]
