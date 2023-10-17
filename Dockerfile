FROM golang:alpine3.18

WORKDIR /

COPY . .

RUN go mod download

RUN go build -o /golang-truewallet-api

EXPOSE 1500

CMD ["/golang-truewallet-api"]