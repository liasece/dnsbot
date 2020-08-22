
all: build

build:
	go build -o bin/dnsbot .

docker-image: build
	docker build -t	liasece/dnsbot:v1.0 .

.PHONY: all build
