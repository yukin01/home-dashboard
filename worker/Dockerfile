FROM golang:1.14.2-buster as builder

WORKDIR /go/src/app

COPY go.* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags="-w -s" -o /app


FROM alpine:3.11

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app /app

EXPOSE 8080

CMD ["/app"]
