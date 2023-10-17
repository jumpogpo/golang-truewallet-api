FROM golang:latest

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /golang-truewallet-api

EXPOSE 1500

CMD ["/golang-truewallet-api"]