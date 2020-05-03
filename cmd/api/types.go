package api

import (
	"github.com/coreos/go-systemd/v22/dbus"
)

// jsonResp is a struct for hanlding json reponse
type jsonResp struct {
	StatusCode int    `json:"code"`
	UnitName   string `json:"unitname"`
	UnitStatus string `json:"unitstatus"`
}

// conf struct handles configuration for systemd service
type conf struct {
	Target     string
	Connection *dbus.Conn
}
