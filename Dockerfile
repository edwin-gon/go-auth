FROM golang:1.24

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN go build -v -o /usr/local/bin/main ./...

ENTRYPOINT [ "main" ]