FROM golang:1.17-alpine3.14
# RUN apk add build-base
WORKDIR /ln-backend

COPY . .
RUN go mod download
RUN go build -o main .
EXPOSE 8888

CMD ["./main"]