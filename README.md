# evdns
dns api to different dns providers

## license
BSD-1-clause

## install
```
go get evalgo.org/evdns/cmd/evdns
```

## Hetzner

### features implemented
- zones
- [x] display zone id and name
- [x] display most zone info
- [x] display a zone with a given zone id
- [x] create zone
- [x] delete zone
- [x] update a zone
- [x] validate zone file plain
- [x] import zone file plain
- [x] export zone file plain
- records
- [x] display records for a given zone
- [x] display a record with a given record id
- [x] create record for a given zone
- [x] delete a record with a give record id
- [x] update a record
- [x] bulk create records
- [x] bulk update records

### hetzner cli usage
```
# first create a evdns.json configuration file
echo '{"url":"https://dns.hetzner.com/api/v1","token":"YOUR-HETZNER-DNS-TOKEN"}' > evdns.json

# display all zones
evdns hetzner -z

# display a zone
evdns hetzner --zone --id "YOUR-ZONE-ID"

# create a zone
evdns hetzner --create --zone --name "YOUR-ZONE-NAME"

# delete a zone
evdns hetzner --delete --zone --id "YOUR-ZONE-ID"

# update a zone
evdns hetzner --update --zone --id "YOUR-ZONE-ID" --name "YOUR-ZONE-NAME" --ttl "YOUR-TTL"

# export a zone file
evdns hetzner --export --id "YOUR-ZONE-ID"

# validate a zone file
evdns hetzner --validate --value "YOUR-ZONE-ID-FILE-PATH.zone"

# import a zone file
evdns hetzner --import --id "YOUR-ZONE-ID" --value "YOUR-ZONE-ID-FILE-PATH.zone"

# display records for a given zone
evdns hetzner -r --id "YOUR-ZONE-ID"

# display a record with a given id
evdns hetzner --record --id "YOUR-RECORD-ID"

# create a record for a given zone
evdns --create --record --id YOUR-ZONE-ID --type "A" --name "YOUR-SUBDOMAIN" --value "YOUR-IP-ADDRESS"

# delete a record with a given id
evdns --delete --record --id YOUR-RECORD-ID

# update a record for a given zone
evdns --update --record --rid YOUR-RECORD-ID --id YOUR-ZONE-ID --type "A" --name "YOUR-SUBDOMAIN" --value "YOUR-IP-ADDRESS"

# create records
evdns --create --records --value '{"records": [{"value": "YOUR-IP-VALUE","type": "A","name": "YOUR-NAME","zone_id": "YOUR-ZONE-ID"}]}'

# update records
evdns --update --records --value '{"records": [{"id":"YOUR-RECORD-ID","value": "YOUR-IP-VALUE","type": "A","name": "YOUR-NAME","zone_id": "YOUR-ZONE-ID"}]}'
```

### hetzner service usage
```
# first create a evdns-service.json configuration file
echo '{"url":"https://dns.hetzner.com/api/v1","token":"YOUR-HETZNER-DNS-TOKEN","webroot":"./webroot","address":":8888"}' > evdns-service.json

# start the evdns-service websocket server
evdns-service start

# download./install the evmsg-js websocket client

# use the evmsg-js client to display the zones

```
