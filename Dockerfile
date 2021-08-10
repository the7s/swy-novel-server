FROM golang:1.16.7-alpine3.14
WORKDIR /data
COPY . .
RUN go mod tidy
CMD ["go","run","main.go"]