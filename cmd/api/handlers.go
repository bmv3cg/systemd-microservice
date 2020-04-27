package api

import (
	"fmt"
	"net/http"

	sys "github.com/bmv3cg/systemd-microservice/pkg/systemd"
)

func ApiIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
	fmt.Fprint(w, "systemd manager microservice\n")
}

func Health(w http.ResponseWriter, r *http.Request) {
	conn := sys.StartSystemConn()
	target := "syncthing.service"

	fmt.Println(w, sys.ServiceStatus(conn, target))

	if sys.ServiceStatus(conn, target) == true {
		fmt.Fprint(w, "systemd unit ", target, " is running!\n")
	} else {
		fmt.Fprint(w, "systemd unit ", target, " not running\n")
	}
}

func StartUnit(w http.ResponseWriter, r *http.Request) {
	conn := sys.StartSystemConn()
	reschan := make(chan string)
	target := "syncthing.service"
	if sys.StartService(conn, target, reschan) == true {
		fmt.Fprint(w, "systemd unit ", target, " started\n")
	} else {
		fmt.Fprint(w, "systemd unit ", target, " failed to start\n")
	}
}

func StopUnit(w http.ResponseWriter, r *http.Request) {
	conn := sys.StartSystemConn()
	reschan := make(chan string)
	target := "syncthing.service"
	if sys.StopService(conn, target, reschan) == true {
		fmt.Fprint(w, "systemd unit ", target, " stopped\n")
	} else {
		fmt.Fprint(w, "systemd unit ", target, " failed to stop\n")
	}
}

func RestartUnit(w http.ResponseWriter, r *http.Request) {
	conn := sys.StartSystemConn()
	reschan := make(chan string)
	target := "syncthing.service"
	if sys.StartService(conn, target, reschan) == true {
		fmt.Fprint(w, "systemd unit ", target, " restarted\n")
	} else {
		fmt.Fprint(w, "systemd unit ", target, " failed to restart\n")
	}
}
