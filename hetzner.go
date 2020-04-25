package evdns

import (
	"bytes"
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Hetzner defines the Hetzner struct
type Hetzner struct {
	// ApiURL where to send the request to
	ApiURL string
	// Token is needed for authorization
	Token string
}

// NewHetzner returns a Hetzner object witht the given api url and token
func NewHetzner(apiURL, token string) *Hetzner {
	return &Hetzner{ApiURL: apiURL, Token: token}
}

// dnsRequest executes the request and returns the response
func dnsRequest(hd *Hetzner, rType string, values url.Values) ([]byte, error) {
	reqURL := hd.ApiURL
	method := "GET"
	body := []byte("")
	switch rType {
	case "zones":
		reqURL += "/zones"
		body = nil
	case "zoneByID":
		reqURL += "/zones/" + values.Get("id")
		body = nil
	case "deleteZone":
		method = "DELETE"
		reqURL += "/zones/" + values.Get("id")
		body = nil
	case "newZone":
		reqURL += "/zones"
		method = "POST"
		body = []byte(values.Get("body"))
	case "records":
		reqURL += "/records?" + values.Encode()
		body = nil
	case "recordByID":
		reqURL += "/records/" + values.Get("id")
		body = nil
	case "deleteRecord":
		method = "DELETE"
		reqURL += "/records/" + values.Get("id")
		body = nil
	case "newRecord":
		reqURL += "/records"
		method = "POST"
		body = []byte(values.Get("body"))
	default:
		body = []byte(values.Encode())
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, reqURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	if rType == "newRecord" {
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Add("Auth-API-Token", hd.Token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// Zones returns the zones available in the Hetzner dns system
func (hd *Hetzner) Zones() (interface{}, error) {
	zonesJSON, err := dnsRequest(hd, "zones", nil)
	if err != nil {
		return nil, err
	}
	var zones interface{}
	err = json.Unmarshal(zonesJSON, &zones)
	if err != nil {
		return nil, err
	}
	return zones, nil
}

// NewZone creates a new zone
func (hd *Hetzner) NewZone(zone map[string]interface{}) (interface{}, error) {
	bodyJSON, err := json.Marshal(zone)
	if err != nil {
		return nil, err
	}
	values := url.Values{"body": []string{string(bodyJSON)}}
	createJSON, err := dnsRequest(hd, "newZone", values)
	if err != nil {
		return nil, err
	}
	var create interface{}
	err = json.Unmarshal(createJSON, &create)
	if err != nil {
		return nil, err
	}
	return create, nil
}

// Zone returns a zone with the given id
func (hd *Hetzner) Zone(zoneID string) (interface{}, error) {
	zonesJSON, err := dnsRequest(hd, "zoneByID", url.Values{"id": []string{zoneID}})
	if err != nil {
		return nil, err
	}
	var zone interface{}
	err = json.Unmarshal(zonesJSON, &zone)
	if err != nil {
		return nil, err
	}
	return zone, nil
}

// DeleteZone deletes a zone with the given zoneID
func (hd *Hetzner) DeleteZone(zoneID string) (interface{}, error) {
	deleteJSON, err := dnsRequest(hd, "deleteZone", url.Values{"id": []string{zoneID}})
	if err != nil {
		return nil, err
	}
	var zoneDeleted interface{}
	err = json.Unmarshal(deleteJSON, &zoneDeleted)
	if err != nil {
		return nil, err
	}
	return zoneDeleted, nil
}

// NewRecord cretes a new record
func (hd *Hetzner) NewRecord(record map[string]interface{}) (interface{}, error) {
	bodyJSON, err := json.Marshal(record)
	if err != nil {
		return nil, err
	}
	values := url.Values{"body": []string{string(bodyJSON)}}
	createJSON, err := dnsRequest(hd, "newRecord", values)
	if err != nil {
		return nil, err
	}
	var create interface{}
	err = json.Unmarshal(createJSON, &create)
	if err != nil {
		return nil, err
	}
	return create, nil
}

// Records returns all records for a given zoneID
func (hd *Hetzner) Records(zoneID string) (interface{}, error) {
	recordsJSON, err := dnsRequest(hd, "records", url.Values{"zone_id": []string{zoneID}})
	if err != nil {
		return nil, err
	}
	var records interface{}
	err = json.Unmarshal(recordsJSON, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// Record returns a record with a given recordID
func (hd *Hetzner) Record(recordID string) (interface{}, error) {
	recordJSON, err := dnsRequest(hd, "recordByID", url.Values{"id": []string{recordID}})
	if err != nil {
		return nil, err
	}
	var record interface{}
	err = json.Unmarshal(recordJSON, &record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// DeleteRecord deletes a record with the given recordID
func (hd *Hetzner) DeleteRecord(recordID string) (interface{}, error) {
	deleteJSON, err := dnsRequest(hd, "deleteRecord", url.Values{"id": []string{recordID}})
	if err != nil {
		return nil, err
	}
	var deleted interface{}
	err = json.Unmarshal(deleteJSON, &deleted)
	if err != nil {
		return nil, err
	}
	return deleted, nil
}
