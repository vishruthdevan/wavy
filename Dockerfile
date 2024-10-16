FROM golang:1.23.2-alpine

WORKDIR /wavy

COPY go.mod ./

RUN go mod download

COPY . .

WORKDIR /wavy/lexer

ENTRYPOINT ["go", "test", "-v", "./...", "-args"]
