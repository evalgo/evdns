# evdns
dns api to different dns providers

## license
BSD-1-clause

## install
```
go get evalgo.org/evdns/cmd/evdns
```

## Hetzner

### hetzner usage
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

```

### features implemented
- zones
- [x] display zone id and name
- [x] display most zone info
- [x] display a zone with a given zone id
- [x] create zone
- [x] delete zone
- [x] update a zone
- [ ] validate zone file plain
- [ ] import zone file plain
- [ ] export zone file plain
- records
- [x] display records for a given zone
- [x] display a record with a given record id
- [x] create record for a given zone
- [x] delete a record with a give record id
- [x] update a record
- [ ] bulk create records
- [ ] bulk update records
