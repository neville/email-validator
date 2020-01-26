FROM golang:1.13.6

ADD . /go/src/email-validation-service

WORKDIR /go/src/email-validation-service/src

RUN go build -o "email-validation-service"

CMD ["./email-validation-service"]

EXPOSE 8080