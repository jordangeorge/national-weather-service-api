FROM golang:1.23.5

WORKDIR /app

# for caching
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o api

EXPOSE 8080

CMD ["./api"]