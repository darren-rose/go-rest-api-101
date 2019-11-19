FROM golang:1.12

LABEL maintainer="Darren Rose <darrenwrose@gmail.com>"

WORKDIR /app

COPY . /app

RUN go build -o app cmd/main.go

CMD ["./app"]

