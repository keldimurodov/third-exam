FROM golang:1.22.1-alpine3.19

RUN mkdir app

COPY . /app

WORKDIR /app

RUN go build -o main cmd/main.go

CMD ["/app/main"]

EXPOSE 9000