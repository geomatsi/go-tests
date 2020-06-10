package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	dbus "github.com/godbus/dbus/v5"
	introspect "github.com/godbus/dbus/v5/introspect"
)

func main() {
	conn, err := dbus.SystemBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to system bus:", err)
		os.Exit(1)
	}

	defer conn.Close()

	args := os.Args
	if len(args) < 2 {
		help(args)
		os.Exit(1)
	}

	switch args[1] {
	case "modems":
		dbusListModems(conn)
	case "modem":
		if len(args) < 4 {
			help(args)
			os.Exit(1)
		}
		dbusModem(conn, args[2], args[3])
	case "intro":
		if len(args) < 4 {
			help(args)
			os.Exit(1)
		}
		dbusIntro(conn, args[2], args[3])
	default:
		help(args)
		os.Exit(1)
	}
}

func dbusIntro(conn *dbus.Conn, service string, path string) {
	fmt.Printf("Introspect %s%s:\n", service, path)
	node, err := introspect.Call(conn.Object(service, dbus.ObjectPath(path)))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to perform introspection:", err)
		return
	}

	data, _ := json.MarshalIndent(node, "", "    ")
	os.Stdout.Write(data)
	fmt.Println("")
}

func dbusListModems(conn *dbus.Conn) {
	var modems [][]interface{}

	// org.ofono.Manager.GetModems: a(oa{sv})
	err := conn.Object("org.ofono", dbus.ObjectPath("/")).Call("org.ofono.Manager.GetModems", 0).Store(&modems)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get available modems:", err)
		return
	}

	for _, mx := range modems {
		switch reflect.TypeOf(mx).Kind() {
		case reflect.Slice:
			for _, m := range mx {
				switch m.(type) {
				case dbus.ObjectPath:
					fmt.Printf("modem %s\n", m)
				}
			}
		}
	}
}

func dbusModem(conn *dbus.Conn, cmd string, modem string) {
	// DBus type: a{sv}
	var res map[string]dbus.Variant
	var method string

	switch cmd {
	case "dump":
		method = "org.ofono.Modem.GetProperties"
		fmt.Printf("Modem %s properties:\n", modem)
	case "sim":
		method = "org.ofono.SimManager.GetProperties"
		fmt.Printf("Modem %s SIM properties:\n", modem)
	case "stats":
		method = "org.ofono.cinterion.HardwareMonitor.GetStatistics"
		fmt.Printf("Modem %s hardware stats:\n", modem)
	case "aps":
		method = "org.ofono.AllowedAccessPoints.GetAllowedAccessPoints"
		fmt.Printf("Modem %s APs:\n", modem)
	case "conn":
		method = "org.ofono.ConnectionManager.GetProperties"
		fmt.Printf("Modem %s connection properties:\n", modem)
	case "reg":
		method = "org.ofono.NetworkRegistration.GetProperties"
		fmt.Printf("Modem %s connection properties:\n", modem)
	default:
		fmt.Printf("Unknown modem command: %s\n", cmd)
		return
	}

	err := conn.Object("org.ofono", dbus.ObjectPath(modem)).Call(method, 0).Store(&res)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to call DBus method:", err)
		return
	}

	fmt.Println("{")

	for k, v := range res {
		fmt.Printf("\t%s: %v\n", k, v)
	}

	fmt.Println("}")
	return
}

func help(args []string) {
	fmt.Printf("%s <command> <param1> <param2> ...\n", args[0])
	fmt.Printf("\t%s intro <service> <object> - DBus object introspection\n", args[0])
	fmt.Printf("\t%s modems - list available oFono modems\n", args[0])
	fmt.Printf("\t%s modem dump <name> - list modem properties\n", args[0])
	fmt.Printf("\t%s modem sim <name> - list modem SIM properties\n", args[0])
	fmt.Printf("\t%s modem stats <name> - list modem stats (Cinterion modems only)\n", args[0])
	fmt.Printf("\t%s modem aps <name> - list modem Access Points\n", args[0])
	fmt.Printf("\t%s modem conn <name> - list modem connection properties\n", args[0])
	fmt.Printf("\t%s modem reg <name> - list modem network registration properties\n", args[0])
}
