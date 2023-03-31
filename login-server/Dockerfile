FROM golang:alpine

WORKDIR /app

COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct&&go build -o main .

EXPOSE 8080

CMD ["./main"]
