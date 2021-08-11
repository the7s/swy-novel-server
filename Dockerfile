FROM golang:1.16.7-alpine3.14
WORKDIR /data
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
CMD export GIN_MODE=release
CMD ["go","run","main.go"]