FROM gcr.io/distroless/base
COPY app .
CMD ["./app"]
EXPOSE 8080