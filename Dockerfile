FROM golang:1.22.1
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest \
    && go install github.com/spf13/cobra-cli@latest

COPY go.mod go.sum ./
RUN go mod download
COPY . .
CMD ["air", "-c", ".air.toml"]
