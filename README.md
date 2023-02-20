Automatically listens to your local IP and updates your DNS record on Namesilo etc.

# Run with shell

```shell
go run . -s namesilo -k "your api key" -i 20s -d example.com -d 1.example.com -d 2.example.com -p http://ifconfig.me
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
docker run --name myDNSBot --restart always -v /home/docker/dnsbot/logs:/dnsbot/logs -d liasece/dnsbot -s namesilo -k "your api key" -i 20s -d example.com -d 1.example.com -d 2.example.com -p http://ifconfig.me
```

Usage of dnsbot:

```shell
  -d value
        The domain associated list with the DNS resource record to modify
  -i duration
        The time of updating interval, like: 30s or 5m (default 5m0s)
  -k string
         To be replaced by your unique API key. Visit the API Manager page within your account for details. (default "***your key***")
  -p string
        The get host external ip service host. (default "http://ifconfig.me")
  -s string
        The dns service name. (default "namesilo")
```
