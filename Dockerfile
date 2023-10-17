FROM golang:alpine3.18

WORKDIR /

ENV GIN_MODE=release

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /golang-truewallet-api

EXPOSE 1500

CMD ["/golang-truewallet-api"]