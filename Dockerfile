FROM golang:1.24.0-bookworm

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /myapp

CMD ["/myapp"]
