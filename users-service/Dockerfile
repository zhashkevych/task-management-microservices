FROM golang:latest

RUN go version

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o users-service ./cmd/main.go

CMD ["./users-service"]