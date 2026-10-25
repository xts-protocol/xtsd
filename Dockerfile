FROM golang:1.27

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o xtsd ./cmd/xtsd

EXPOSE 8080

CMD ["./xtsd"]
