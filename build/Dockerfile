# Using the multi stage build process to reduce image size
# Build stage
FROM golang:1.13.6-alpine AS builder

ADD ./ /go/src/email-validator

WORKDIR /go/src/email-validator/cmd
 
RUN go build -o "email-validator"

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /go/src/email-validator/cmd/email-validator .

CMD ["./email-validator"]

EXPOSE 8080