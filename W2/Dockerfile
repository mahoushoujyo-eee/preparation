# ===== Stage 1: Build =====
FROM golang:1.24-alpine AS builder

# 国内构建时可用来加速
# ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /src

# 先拷贝 module 文件，利用 layer 缓存（需要学习的点）
COPY go.mod go.sum ./
RUN go mod download

# 拷贝源码并编译
COPY . .

# CGO_ENABLED=0 生成静态二进制，便于 alpine/scratch 运行
# -trimpath 去掉本地路径信息，减小体积并避免泄露
# -ldflags="-s -w" 去掉符号表与 DWARF 调试信息
# -buildvcs=false 避免在容器里嵌入 git 信息失败
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -buildvcs=false \
      -ldflags="-s -w" \
      -o /out/app .

# ===== Stage 2: Runtime =====
FROM alpine:3.20

# 必要的运行时依赖：CA 证书（用于 HTTPS / 内部服务调用）、时区数据、非 root 用户
RUN apk add --no-cache ca-certificates tzdata && \
    addgroup -S app && adduser -S app -G app

WORKDIR /app

# 从构建阶段拷贝产物
COPY --from=builder /out/app /app/app

# 服务监听端口（与代码中 127.0.0.1:8085 一致，容器内需要暴露）
EXPOSE 8085

USER app

ENTRYPOINT ["/app/app"]