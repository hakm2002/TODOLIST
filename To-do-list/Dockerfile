# Stage 1: Build
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Copy dependency dulu (Cache Layering)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Run
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

# Install Timezone & Certificates (Penting untuk API Call HTTPS)
RUN apk add --no-cache tzdata ca-certificates
ENV TZ=Asia/Jakarta

EXPOSE 8080
CMD ["./main"]
