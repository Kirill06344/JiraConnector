FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/main

RUN CGO_ENABLED=0 GOOS=linux go build -o /connector

WORKDIR /app

CMD ["/connector"]