FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o api-gateway .

EXPOSE 8002

COPY ./tls/cert.pem /tls/cert.pem
COPY ./tls/key.pem /tls/key.pem

CMD ["./api-gateway"]