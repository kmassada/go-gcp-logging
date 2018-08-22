FROM golang:latest as builder
RUN apt-get install git
COPY main.go .
RUN go get -u cloud.google.com/go/logging
RUN go build -o /app main.go

FROM alpine:latest 
CMD ["./app"]
COPY --from=builder /app .
EXPOSE 8080