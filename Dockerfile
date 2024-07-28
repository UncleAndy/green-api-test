FROM golang:1.22.4 AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN CGO_ENABLED=0 go build -o green-api-test ./cmd/main.go

FROM alpine

COPY --from=builder /build/green-api-test /green-api-test
EXPOSE 8080

CMD ["/green-api-test"]
