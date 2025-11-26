# ---------- Stage 1: Build ----------
FROM golang:1.24.2 AS builder

# Set workdir
WORKDIR /app

# Copy mod files first (agar cache efektif)
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary untuk Linux (Docker container)
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -ldflags="-s -w" -o main .

# ---------- Stage 2: Run ----------
FROM alpine:latest

WORKDIR /app

# Copy hasil build dari stage builder
COPY --from=builder /app/main .

# Expose port default Gin
EXPOSE 9909

# Start the app
CMD ["./main"]