FROM golang:1.23.1 as builder
# Set the application directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the application source
COPY . .
# Build the application
RUN CGO_ENABLED=0 go build -o main cmd/api/main.go

# Execution stage
FROM gcr.io/distroless/base-debian10

# Copy the built binary
COPY --from=builder /app/main /
COPY  pack_sizes.json /
COPy .env /
# Execute the application
CMD ["/main"]