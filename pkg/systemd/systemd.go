package systemd

import (
	"fmt"
	"log"

	"github.com/coreos/go-systemd/v22/dbus"
)

// StartSystemConn is a function to initalise connection systemd as user.
// Systemd connection is made as user with minimal privllages.
func StartSystemConn() (conn *dbus.Conn) {
	conn, err := dbus.NewUserConnection()
	if err != nil {
		log.Fatalf("Cannot create Systemd connection: %v\n", err)
	}
	return conn
}

// StartSystemConn is a function to initalise connection systemd without using Dbus
func SystemdConn() (conn *dbus.Conn) {
	conn, err := dbus.NewSystemdConnection()
	if err != nil {
		log.Fatalf("Cannot create Systemd connection: %v\n", err)
	}
	return conn
}

// ListUnits list all systemd service available
func ListUnits(conn *dbus.Conn) {
	units, err := conn.ListUnits()
	if err != nil {
		log.Fatalf("can't list units: %v", err)
	}
	log.Println("Loaded units: %+v", units)

}

// ServiceStatus return status of systemd service
func ServiceStatus(conn *dbus.Conn, target string) (status bool) {
	unitState, err := conn.GetUnitProperty(target, "ActiveState")
	if err != nil {
		log.Fatalln("Cant get service state")
	}
	if unitState.Value.String() == "\"active\"" {
		log.Print(target, " is alive")
		return true
	} else {
		log.Print(target, " is not alive")
		return false
	}
}

// StopService stops systemd unit
func StopService(conn *dbus.Conn, target string, reschan chan string) (re bool) {
	_, err := conn.StopUnit(target, "replace", reschan)
	if err != nil {
		fmt.Println(err)
	}
	status := <-reschan
	if status != "done" {
		return false
	} else {
		return true
	}

}

// RestartService restarts systemd unit
func RestartService(conn *dbus.Conn, target string, reschan chan string) (re bool) {
	_, err := conn.RestartUnit(target, "replace", reschan)
	if err != nil {
		fmt.Println(err)
	}
	status := <-reschan
	if status != "done" {
		return false
	} else {
		return true
	}

}

// StartService start's systemd unit
func StartService(conn *dbus.Conn, target string, reschan chan string) (re bool) {
	_, err := conn.StartUnit(target, "replace", reschan)
	if err != nil {
		fmt.Println(err)
	}
	status := <-reschan
	if status != "done" {
		return false
	} else {
		return true
	}
}
