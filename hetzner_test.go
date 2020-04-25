package evdns

import (
	"testing"
)

func Test_Unit_Hetzner(t *testing.T) {
	h := NewHetzner("api", "token")
	if h.ApiURL != "api" || h.Token != "token" {
		t.Error("could not create a Hetzner object")
	}
}
