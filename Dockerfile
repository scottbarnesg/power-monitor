FROM golang:1.22

WORKDIR /app
# Download and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy local .go files and compile them
COPY ./ ./
RUN go build
CMD ["./power-monitor", "-server"]