package api

import (
	"fmt"
	"net/http"

	sys "github.com/bmv3cg/systemd-microservice/pkg/systemd"
	"github.com/spf13/viper"
)

// ApIIndex function for servinf API index 
func ApIIndex(w http.ResponseWriter, r *http.Request) {
	t := "Systemd manager microservice\n"
	//fmt.Fprint(w, "Systemd manager microservice\n")
	j, _ := json.Marshal(t)
	w.Write(j)
}

// Health function check whether systemd fucntion is running 
func Health(w http.ResponseWriter, r *http.Request) {
	conn := sys.StartSystemConn()
	fmt.Println(viper.GetString("systemdunit"))
	target := viper.GetString("systemdunit")
	if sys.ServiceStatus(conn, target) == true {
		fmt.Fprint(w, "systemd unit ", target, " is running!\n")
	} else {
		fmt.Fprint(w, "systemd unit ", target, " not running\n")
	}
}

// StartUnit function starts systemd service which is configured in conf.yaml
func StartUnit(w http.ResponseWriter, r *http.Request) {
	conn := sys.StartSystemConn()
	reschan := make(chan string)
	target := viper.GetString("systemdunit")
	if sys.StartService(conn, target, reschan) == true {
		fmt.Fprint(w, "systemd unit ", target, " started\n")
	} else {
		fmt.Fprint(w, "systemd unit ", target, " failed to start\n")
	}
}

// StopUnit function stops systemd service which is configured in conf.yaml
func StopUnit(w http.ResponseWriter, r *http.Request) {
	conn := sys.StartSystemConn()
	reschan := make(chan string)
	target := viper.GetString("systemdunit")
	if sys.StopService(conn, target, reschan) == true {
		fmt.Fprint(w, "systemd unit ", target, " stopped\n")
	} else {
		fmt.Fprint(w, "systemd unit ", target, " failed to stop\n")
	}
}

// RestartUnit function restarts systemd service which is configured in conf.yaml
func RestartUnit(w http.ResponseWriter, r *http.Request) {
	conn := sys.StartSystemConn()
	reschan := make(chan string)
	target := viper.GetString("systemdunit")
	if sys.StartService(conn, target, reschan) == true {
		fmt.Fprint(w, "systemd unit ", target, " restarted\n")
	} else {
		fmt.Fprint(w, "systemd unit ", target, " failed to restart\n")
	}
}
