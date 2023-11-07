FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o ./bin/healthcare

FROM alpine
WORKDIR /app
COPY --from=builder /app/bin/healthcare .
ENV DB_USERNAME= \
    DB_PASSWORD= \
    DB_HOST= \
    DB_PORT= \
    DB_NAME= \
    JWT_SECRET= 
ENTRYPOINT ["./healthcare"]