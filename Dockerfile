FROM golang:1.12

LABEL maintainer="Darren Rose <darrenwrose@gmail.com>"

WORKDIR /app

COPY go.mod go.sum main.go ./

RUN go build -o app .

CMD ["./app"]

