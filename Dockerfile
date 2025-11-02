FROM golang:1.14.3-bullseye

WORKDIR /dnsbot
COPY ./ ./

ENTRYPOINT ["./bin/dnsbot"]
