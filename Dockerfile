FROM golang:1.23 as builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o his-middleware .

FROM alpine:latest
# RUN apk --no-cache add ca-certificates

RUN apk add --no-cache tzdata ca-certificates && \
    cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && \
    echo "Asia/Bangkok" > /etc/timezone && \
    apk del tzdata

# Copy init script
COPY docker/postgres/init-multiple-dbs.sh /docker-entrypoint-initdb.d/

# Set execute permission
RUN chmod +x /docker-entrypoint-initdb.d/init-multiple-dbs.sh

WORKDIR /root/
COPY --from=builder /app/his-middleware .
COPY --from=builder /app/.env .

EXPOSE 8080
CMD ["./his-middleware"]