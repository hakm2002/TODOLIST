# Stage 1: Build
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Run
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
# Install Timezone (Opsional, biar log waktunya WIB)
RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

EXPOSE 8080
CMD ["./main"]
