FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

COPY wait-for.sh /wait-for.sh
RUN chmod +x /wait-for.sh

RUN go mod tidy
RUN go build -o myapp

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache postgresql-client


COPY --from=builder /app/myapp ./myapp
COPY --from=builder /wait-for.sh /wait-for.sh
COPY .env .env

RUN chmod +x /wait-for.sh

EXPOSE 8080

CMD ["/wait-for.sh", "postgres", "./myapp"]
