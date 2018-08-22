FROM alpine:latest 
COPY app .
CMD ["./app"]
EXPOSE 8080