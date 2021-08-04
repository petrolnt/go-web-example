FROM golang:1.16-alpine AS build
WORKDIR /app
COPY go.mod ./  
RUN go mod download
COPY *.go ./
RUN go build -o docker-web-example

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /app/docker-web-example ./
EXPOSE 8080
CMD [ "./docker-web-example" ]

