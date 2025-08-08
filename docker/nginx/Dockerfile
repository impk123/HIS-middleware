FROM golang:1.20 as builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o HIS-middleware .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/HIS-middleware .
COPY --from=builder /app/.env .

EXPOSE 8080
CMD ["./HIS-middleware"]