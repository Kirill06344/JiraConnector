FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/main

EXPOSE 8000

RUN CGO_ENABLED=0 GOOS=linux go build -o /backend

WORKDIR /app

RUN ls

CMD ["/backend"]