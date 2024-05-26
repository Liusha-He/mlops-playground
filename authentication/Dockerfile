FROM golang:1.21-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app
RUN CGO_ENABLED=0 go build -o userService ./src

RUN chmod +x /app/userService

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/userService /app

CMD ["/app/userService"]
