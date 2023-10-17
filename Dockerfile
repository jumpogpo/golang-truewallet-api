FROM golang:alpine3.18

WORKDIR /

COPY . .

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN go mod download

RUN go build -o /golang-truewallet-api

EXPOSE 1500

CMD ["/golang-truewallet-api"]