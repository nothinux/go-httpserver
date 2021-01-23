FROM golang:1.15.7-alpine3.13 as builder
WORKDIR /go/src/github.com/nothinux/karsajobs
ENV GO111MODULE=on
COPY . .
RUN go mod download
RUN mkdir /build; \
    go build -o /build/ ./...


FROM alpine:latest
WORKDIR /app
RUN mkdir /static
COPY --from=builder /build/go-httpserver .
EXPOSE 8080
CMD ["./app/go-httpserver -workdir /static"]