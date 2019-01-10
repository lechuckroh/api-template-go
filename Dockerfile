FROM alpine:latest

LABEL MAINTAINER="Lechuck Roh <lechuckroh@gmail.com>"

WORKDIR /app
RUN mkdir -p /app

RUN apk add --update curl && rm -rf /var/cache/apk/*

VOLUME ["/app/logs"]

EXPOSE 8080

HEALTHCHECK CMD /app/healthcheck.sh

COPY healthcheck.sh ./
COPY app-server ./

ENTRYPOINT ["./app-server"]
