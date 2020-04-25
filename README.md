# evdns
dns api to different dns providers

## license
BSD-1-clause

## install
```
go get evalgo.org/evdns/cmd/evdns
```

## hetzner usage
```
# first create a evdns.json configuration file
echo '{"url":"https://dns.hetzner.com/api/v1","token":"YOUR-HETZNER-DNS-TOKEN"}' > evdns.json

# run the evdns executable to display the zones
evdns hetzner -z
```
