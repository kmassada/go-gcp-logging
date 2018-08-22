FROM golang:1.10.1-alpine3.7 as builder
COPY main.go .
RUN go get -u cloud.google.com/go/logging
RUN go build -o /app main.go

FROM alpine:3.7  
CMD ["./app"]
COPY --from=builder /app .
EXPOSE 8080