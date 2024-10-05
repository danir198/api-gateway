FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o api-gateway .

EXPOSE 8002

CMD ["./api-gateway"]