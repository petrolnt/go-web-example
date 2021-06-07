FROM golang:latest
COPY . .
RUN go test && go build ./src/main.go

