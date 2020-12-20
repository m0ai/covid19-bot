FROM golang:alpine
ENV GO11MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main .

WORKDIR /dist
COPY .env .
RUN cp /build/main .
CMD ["/dist/main"]
