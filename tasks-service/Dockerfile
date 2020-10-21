FROM golang:latest

RUN go version

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o tasks-service ./cmd/main.go

CMD ["./tasks-service"]