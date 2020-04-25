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

# run the evdns executable to display the zones
evdns hetzner -z
```

### features implemented
- zones
- [x] display zone id and name
- [ ] display all zone info
- [ ] display a zone with a given zone id
- [ ] create zone
- [ ] delete zone
- [ ] update a zone
- [ ] validate zone file plain
- [ ] import zone file plain
- [ ] export zone file plain
- records
- [ ] display records for a given zone
- [ ] display a record with a given record id
- [ ] create record for a given zone
- [ ] delete a record with a give record id
- [ ] update a record
- [ ] bulk create records
- [ ] bulk update records
