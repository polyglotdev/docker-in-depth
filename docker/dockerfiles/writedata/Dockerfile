FROM golang:1.22.4-bookworm

WORKDIR /app

# Copy only the go.mod and go.sum files first
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o writedata .

CMD ["./writedata", "/logs/myapp"]