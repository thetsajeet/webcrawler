FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

ADD webcrawler /usr/bin/webcrawler

ENTRYPOINT [ "webcrawler" ]