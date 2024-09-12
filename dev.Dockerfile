FROM golang:1.22.6-alpine AS builder
WORKDIR /go/bin
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o tyk-gateway

FROM alpine:3.7
COPY --from=builder /go/bin/tyk-gateway /usr/local/bin/tyk-gateway
COPY --from=builder /go/bin/templates /usr/local/bin/templates
WORKDIR /usr/local/bin
RUN chmod +x /usr/local/bin/tyk-gateway
EXPOSE 8080
CMD sh -c "ls -la /usr/local/bin && ./tyk-gateway"