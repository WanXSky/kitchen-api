FROM golang:1.22-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /app

COPY go.mod go.sum ./ 
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go 

FROM alpine:latest
WORKDIR /root

COPY --from=builder /app/main .

# copy Database local jika perlu
# COPY --from=builder /app/kitchen.db .

EXPOSE 3000

CMD ["./main"]
