FROM golang:1.17-alpine

WORKDIR /app

COPY ./src/go.mod ./src/go.sum ./

RUN go mod download

COPY ./src .

RUN go build -o api-server

EXPOSE 2000

CMD ["./api-server"]