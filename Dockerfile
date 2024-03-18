FROM golang:1.22

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go get -u github.com/pressly/goose/cmd/goose

COPY . .
# RUN goose -dir db/migrations postgres "postgres://admin:admin123@127.0.0.1:5432/postgres?sslmode=disable" up

RUN go build -o app .
CMD ["./app"]