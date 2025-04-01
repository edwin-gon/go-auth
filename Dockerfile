FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN go build -v -o /usr/local/bin/ ./...

ENTRYPOINT [ "go-auth" ]