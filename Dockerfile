FROM alpine:latest 
CMD ["./app"]
COPY app .
EXPOSE 8080