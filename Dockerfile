FROM golang:1.15-alpine AS builder
ENV GO11MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY src/go.mod src/go.sum /build/
RUN go mod download
COPY src/ .

RUN go build -a -ldflags '-s' -o main main.go

FROM scratch
WORKDIR /dist
COPY --from=builder /build/main /dist
CMD ["/dist/main"]
