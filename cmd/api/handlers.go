package api

import (
	"fmt"
	"net/http"
	"encoding/json"

	sys "github.com/bmv3cg/systemd-microservice/pkg/systemd"
	"github.com/spf13/viper"
)
type jsonErr struct {
	StatusCode int `json:"code"`
	UnitName string `json:"unitname"`
    UnitStatus string `json:"unitstatus"`
}

// ApIIndex function for serving API index 
func ApIIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Systemd manager microservice\n")
}

// Health function check whether systemd unit is running 
func Health(w http.ResponseWriter, r *http.Request) {
	conn := sys.StartSystemConn()
	target := viper.GetString("systemdunit")
	if sys.ServiceStatus(conn, target) == true {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonErr{StatusCode: http.StatusOK, UnitName: target, UnitStatus: "systemd unit running"}); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(jsonErr{StatusCode: http.StatusInternalServerError, UnitName: target, UnitStatus: "systemd unit is not running"}); err != nil {
			panic(err)
		}
	}
}

// StartUnit function starts systemd service which is configured in conf.yaml
func StartUnit(w http.ResponseWriter, r *http.Request) {
	conn := sys.StartSystemConn()
	reschan := make(chan string)
	target := viper.GetString("systemdunit")
	if sys.StartService(conn, target, reschan) == true {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonErr{StatusCode: http.StatusOK, UnitName: target, UnitStatus: "systemd unit started"}); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(jsonErr{StatusCode: http.StatusInternalServerError, UnitName: target, UnitStatus: "failed to start systemd unit"}); err != nil {
			panic(err)
		}
	}
}

// StopUnit function stops systemd service which is configured in conf.yaml
func StopUnit(w http.ResponseWriter, r *http.Request) {
	conn := sys.StartSystemConn()
	reschan := make(chan string)
	target := viper.GetString("systemdunit")
	if sys.StopService(conn, target, reschan) == true {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonErr{StatusCode: http.StatusOK, UnitName: target, UnitStatus: "systemd unit stopped"}); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(jsonErr{StatusCode: http.StatusInternalServerError, UnitName: target, UnitStatus: "failed to stop systemd unit"}); err != nil {
			panic(err)
		}
	}
}

// RestartUnit function restarts systemd service which is configured in conf.yaml
func RestartUnit(w http.ResponseWriter, r *http.Request) {
	conn := sys.StartSystemConn()
	reschan := make(chan string)
	target := viper.GetString("systemdunit")
	if sys.StartService(conn, target, reschan) == true {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonErr{StatusCode: http.StatusOK, UnitName: target, UnitStatus: "systemd unit restarted"}); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(jsonErr{StatusCode: http.StatusInternalServerError, UnitName: target, UnitStatus: "failed to restart systemd unit"}); err != nil {
			panic(err)
		}
	}
}
