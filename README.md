# Run with shell

```shell
go run . -s namesilo -k "your api key" -i 20s -d example.com -d 1.example.com -d 2.example.com
```

# Run with docker

```shell
make docker-image
docker run --name myDNSBot --restart always -v /home/docker/dnsbot/logs:/dnsbot/logs -d liasece/dnsbot -s namesilo -k "your api key" -i 20s -d example.com -d 1.example.com -d 2.example.com
```
