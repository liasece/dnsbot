FROM golang:1.14.3-buster

WORKDIR /dnsbot
COPY ./ ./

ENTRYPOINT ["./dnsbot/bin/dnsbot"]
