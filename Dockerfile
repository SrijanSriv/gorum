FROM golang:latest

LABEL maintainer="@srijansriv <srijansrivastav01@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8081

RUN go build .

CMD ["./gorum"]