FROM golang:1.22 as builder
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/client/main.go cmd/client/init.go   

CMD ["./app"]

