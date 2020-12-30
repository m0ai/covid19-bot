FROM golang:alpine as builder
ENV GO11MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY workers .
RUN go build -a -ldflags '-s' -o main main.go

FROM scratch
WORKDIR /dist
COPY --from=builder /build/main /dist
COPY .env .
CMD ["/dist/main"]
