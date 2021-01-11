FROM golang:1.15-alpine AS builder
ENV GO11MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /build
COPY src/go.mod src/go.sum /build/
RUN go mod download
COPY src/ .
# If build is slow, separate  a builder each binary
RUN go build -a -ldflags '-s' -o main main.go
RUN go build -a -ldflags '-s' -o scrapper scrapper.go


FROM scratch AS notify
WORKDIR /dist
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/main /dist/
CMD ["/dist/main"]


FROM scratch AS scrapper
WORKDIR /dist
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/scrapper /dist/
CMD ["/dist/scrapper"]
