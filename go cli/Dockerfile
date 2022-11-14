FROM golang:1.19

COPY . /go-cli-app

WORKDIR /go-cli-app

RUN go build -o main && chmod 777 /go-cli-app

ENTRYPOINT ["/go-cli-app/main"]