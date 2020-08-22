# Run with shell

```shell
go run . -s namesilo -k "your api key" -i 20s -d example.com -d 1.example.com -d 2.example.com
```

# Run with docker

## build form source

```shell
make build
docker build -t liasece/dnsbot .
```

## get image from docker hub

```shell
docker pull liasece/dnsbot
```

## run

```shell
docker run --name myDNSBot --restart always -v /home/docker/dnsbot/logs:/dnsbot/logs -d liasece/dnsbot -s namesilo -k "your api key" -i 20s -d example.com -d 1.example.com -d 2.example.com
```
