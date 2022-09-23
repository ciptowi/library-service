# step 1: build executable binary
FROM golang:1.18-alpine
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /library-service

# step 2: build a small image
FROM alpine:3.16.0
WORKDIR /app
COPY --from=builder library-service .
COPY .env .
EXPOSE 8080
CMD ["./library-service", "-m=migrate"]
