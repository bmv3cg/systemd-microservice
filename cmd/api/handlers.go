package api

import (
	"fmt"
	"net/http"

	sys "github.com/bmv3cg/systemd-microservice/pkg/systemd"
	"github.com/bmv3cg/systemd-microservice/pkg/util"
	"github.com/spf13/viper"
)

// Index function for serving API index
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Systemd manager microservice\n")
}

// Health function check whether systemd unit is running
func Health(w http.ResponseWriter, r *http.Request) {
	c := conf{Target: viper.GetString("systemdunit"), Connection: sys.StartSystemConn()}
	if sys.ServiceStatus(c.Connection, c.Target) == true {
		resp := jsonResp{StatusCode: http.StatusOK, UnitName: c.Target, UnitStatus: "systemd unit running"}
		util.JSONResponse(resp, w)
	} else {
		resp := jsonResp{StatusCode: http.StatusInternalServerError, UnitName: c.Target, UnitStatus: "systemd unit not running"}
		util.JSONResponse(resp, w)
	}
}

// StartUnit function starts systemd service which is configured in conf.yaml
func StartUnit(w http.ResponseWriter, r *http.Request) {
	reschan := make(chan string)
	c := conf{Target: viper.GetString("systemdunit"), Connection: sys.StartSystemConn()}
	if sys.StartService(c.Connection, c.Target, reschan) == true {
		resp := jsonResp{StatusCode: http.StatusOK, UnitName: c.Target, UnitStatus: "systemd unit started"}
		util.JSONResponse(resp, w)
	} else {
		resp := jsonResp{StatusCode: http.StatusInternalServerError, UnitName: c.Target, UnitStatus: "failed to start systemd unit"}
		util.JSONResponse(resp, w)
	}
}

// StopUnit function stops systemd service which is configured in conf.yaml
func StopUnit(w http.ResponseWriter, r *http.Request) {
	reschan := make(chan string)
	c := conf{Target: viper.GetString("systemdunit"), Connection: sys.StartSystemConn()}
	if sys.StopService(c.Connection, c.Target, reschan) == true {
		resp := jsonResp{StatusCode: http.StatusOK, UnitName: c.Target, UnitStatus: "systemd unit stopped"}
		util.JSONResponse(resp, w)
	} else {
		resp := jsonResp{StatusCode: http.StatusInternalServerError, UnitName: c.Target, UnitStatus: "failed to stop systemd unit"}
		util.JSONResponse(resp, w)
	}
}

// RestartUnit function restarts systemd service which is configured in conf.yaml
func RestartUnit(w http.ResponseWriter, r *http.Request) {
	reschan := make(chan string)
	c := conf{Target: viper.GetString("systemdunit"), Connection: sys.StartSystemConn()}
	if sys.StartService(c.Connection, c.Target, reschan) == true {
		resp := jsonResp{StatusCode: http.StatusOK, UnitName: c.Target, UnitStatus: "systemd unit restarted"}
		util.JSONResponse(resp, w)
	} else {
		resp := jsonResp{StatusCode: http.StatusInternalServerError, UnitName: c.Target, UnitStatus: "failed to restart systemd unit"}
		util.JSONResponse(resp, w)
	}
}
