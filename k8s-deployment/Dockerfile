FROM golang:1.19

WORKDIR /k8s-app

COPY . /k8s-app/

RUN go build -o main && chmod 777 /k8s-app

ENTRYPOINT ["./main"]