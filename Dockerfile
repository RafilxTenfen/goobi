FROM alpine:latest
WORKDIR /app

ADD ./goobi /app
CMD ["./goobi", "-l"]
