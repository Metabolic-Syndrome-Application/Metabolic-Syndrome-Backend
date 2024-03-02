FROM golang:alpine3.18

ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download -x

COPY . .

RUN go build -o main .

EXPOSE 8000

CMD ["./main"]