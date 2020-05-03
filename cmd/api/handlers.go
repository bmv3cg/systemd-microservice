package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	sys "github.com/bmv3cg/systemd-microservice/pkg/systemd"
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
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonResp{StatusCode: http.StatusOK, UnitName: c.Target, UnitStatus: "systemd unit running"}); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(jsonResp{StatusCode: http.StatusInternalServerError, UnitName: c.Target, UnitStatus: "systemd unit is not running"}); err != nil {
			panic(err)
		}
	}
}

// StartUnit function starts systemd service which is configured in conf.yaml
func StartUnit(w http.ResponseWriter, r *http.Request) {
	reschan := make(chan string)
	c := conf{Target: viper.GetString("systemdunit"), Connection: sys.StartSystemConn()}
	if sys.StartService(c.Connection, c.Target, reschan) == true {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonResp{StatusCode: http.StatusOK, UnitName: c.Target, UnitStatus: "systemd unit started"}); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(jsonResp{StatusCode: http.StatusInternalServerError, UnitName: c.Target, UnitStatus: "failed to start systemd unit"}); err != nil {
			panic(err)
		}
	}
}

// StopUnit function stops systemd service which is configured in conf.yaml
func StopUnit(w http.ResponseWriter, r *http.Request) {
	reschan := make(chan string)
	c := conf{Target: viper.GetString("systemdunit"), Connection: sys.StartSystemConn()}
	if sys.StopService(c.Connection, c.Target, reschan) == true {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonResp{StatusCode: http.StatusOK, UnitName: c.Target, UnitStatus: "systemd unit stopped"}); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(jsonResp{StatusCode: http.StatusInternalServerError, UnitName: c.Target, UnitStatus: "failed to stop systemd unit"}); err != nil {
			panic(err)
		}
	}
}

// RestartUnit function restarts systemd service which is configured in conf.yaml
func RestartUnit(w http.ResponseWriter, r *http.Request) {
	reschan := make(chan string)
	c := conf{Target: viper.GetString("systemdunit"), Connection: sys.StartSystemConn()}
	if sys.StartService(c.Connection, c.Target, reschan) == true {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonResp{StatusCode: http.StatusOK, UnitName: c.Target, UnitStatus: "systemd unit restarted"}); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(jsonResp{StatusCode: http.StatusInternalServerError, UnitName: c.Target, UnitStatus: "failed to restart systemd unit"}); err != nil {
			panic(err)
		}
	}
}
