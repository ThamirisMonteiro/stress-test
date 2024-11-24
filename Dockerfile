FROM golang:1.19 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o stress-test ./cmd/stresstest

FROM golang:1.19 AS test
WORKDIR /app
COPY --from=builder /app/stress-test .
COPY . .
CMD ["go", "test", "-v", "./..."]

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/stress-test .
RUN chmod +x /root/stress-test
ENTRYPOINT ["/root/stress-test"]
