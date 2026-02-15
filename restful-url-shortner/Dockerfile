# --------------------------------------------------------
# Stage 1: The Builder (Compiles the code)
# --------------------------------------------------------
FROM golang:1.26 AS builder

WORKDIR /app

# Copy dependency files first (this speeds up builds by caching)
COPY go.mod ./
# COPY go.sum ./ # Uncomment if you have this file

# Download modules
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
# CGO_ENABLED=0: Disables C libraries (makes binary portable)
# GOOS=linux: Forces Linux build even if you are on Mac/Windows
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# --------------------------------------------------------
# Stage 2: The Runner (Tiny image to run the app)
# --------------------------------------------------------
FROM alpine:latest

WORKDIR /root/

# Copy only the compiled binary from Stage 1
COPY --from=builder /app/myapp .

# Expose the port so Docker knows we use 8080
EXPOSE 8080

# Command to run when container starts
CMD ["./myapp"]