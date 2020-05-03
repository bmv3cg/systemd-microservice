package api

import (

	"github.com/coreos/go-systemd/v22/dbus"
)

type jsonResp struct {
	StatusCode int    `json:"code"`
	UnitName   string `json:"unitname"`
	UnitStatus string `json:"unitstatus"`
}

type conf struct {
	Target string
	Connection *dbus.Conn
}