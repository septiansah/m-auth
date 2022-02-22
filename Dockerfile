FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app/m-authentication

COPY . .

RUN go mod tidy

RUN go build -o main .

CMD ["./main"]
