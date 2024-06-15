FROM golang:1.22.3-alpine

WORKDIR /go/src/h24s_19

COPY go.mod .
COPY go.sum .

RUN go mod download

CMD ["go", "run", "main.go"]
