FROM golang:1.15.7-alpine3.13 as builder
WORKDIR /go/src/github.com/nothinux/go-httpserver
ENV GO111MODULE=on
COPY . .
RUN go mod download
RUN mkdir /build; \
    go build -o /build/ ./...


FROM alpine:3.13.4
WORKDIR /app
RUN mkdir /static
COPY --from=builder /build/go-httpserver .
EXPOSE 8080
CMD ["/app/go-httpserver", "-dir", "/static"]
