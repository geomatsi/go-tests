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
	case "context":
		if len(args) < 4 {
			help(args)
			os.Exit(1)
		}
		dbusContext(conn, args[2], args[3])

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
	utilSignature1(conn, "/", "org.ofono.Manager.GetModems", false)
}

func dbusContext(conn *dbus.Conn, cmd string, context string) {
	var method string = "org.ofono.ConnectionContext.SetProperty"
	var active dbus.Variant

	switch cmd {
	case "enable":
		active = dbus.MakeVariant(true)
	case "disable":
		active = dbus.MakeVariant(false)
	default:
		fmt.Printf("Unknown context command: %s\n", cmd)
		return
	}

	err := conn.Object("org.ofono", dbus.ObjectPath(context)).Call(method, 0, "Active", active).Store()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to call DBus method:", err)
		return
	}
}

func dbusModem(conn *dbus.Conn, cmd string, modem string) {
	var signature string
	var method string

	switch cmd {
	case "dump":
		method = "org.ofono.Modem.GetProperties"
		signature = "a{sv}"
	case "sim":
		method = "org.ofono.SimManager.GetProperties"
		signature = "a{sv}"
	case "stats":
		method = "org.ofono.cinterion.HardwareMonitor.GetStatistics"
		signature = "a{sv}"
	case "aps":
		method = "org.ofono.AllowedAccessPoints.GetAllowedAccessPoints"
		signature = "as"
	case "conn":
		method = "org.ofono.ConnectionManager.GetProperties"
		signature = "a{sv}"
	case "reg":
		method = "org.ofono.NetworkRegistration.GetProperties"
		signature = "a{sv}"
	case "ctx":
		method = "org.ofono.ConnectionManager.GetContexts"
		signature = "a(oa{sv})"
	case "mon":
		method = "org.ofono.NetworkMonitor.GetServingCellInformation"
		signature = "a{sv}"
	default:
		fmt.Printf("Unknown modem command: %s\n", cmd)
		return
	}

	fmt.Printf("modem[%s] method[%s] signature[%s]:\n",
		modem, method, signature)

	switch signature {
	case "a(oa{sv})":
		utilSignature1(conn, modem, method, true)
	case "a{sv}":
		utilSignature2(conn, modem, method)
	case "as":
		utilSignature3(conn, modem, method)
	}

	return
}

// signature: a(oa{sv})
func utilSignature1(conn *dbus.Conn, path string, method string, dump bool) {
	var res [][]interface{}

	err := conn.Object("org.ofono", dbus.ObjectPath(path)).Call(method, 0).Store(&res)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to call DBus method:", err)
		return
	}

	fmt.Println("{")

	for _, e := range res {
		switch reflect.TypeOf(e).Kind() {
		case reflect.Slice:
			for _, m := range e {
				switch val := m.(type) {
				case dbus.ObjectPath:
					fmt.Printf("\t%s\n", val)
				case map[string]dbus.Variant:
					if dump {
						fmt.Println("\t{")

						for k, v := range val {
							fmt.Printf("\t\t%s: %v\n", k, v)
						}

						fmt.Println("\t}")
					}
				default:
					fmt.Printf("\t%s -> %s\n",
						reflect.TypeOf(m), reflect.TypeOf(m).Kind())
				}
			}
		default:
			fmt.Printf("\t{\n\t\t%s -> %s\n\t}\n",
				reflect.TypeOf(e), reflect.TypeOf(e).Kind())
		}
	}

	fmt.Println("}")
}

// signature: a{sv}
func utilSignature2(conn *dbus.Conn, path string, method string) {
	var res map[string]dbus.Variant

	err := conn.Object("org.ofono", dbus.ObjectPath(path)).Call(method, 0).Store(&res)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to call DBus method:", err)
		return
	}

	fmt.Println("{")

	for k, v := range res {
		fmt.Printf("\t%s: %v\n", k, v)
	}

	fmt.Println("}")
}

// signature: as
func utilSignature3(conn *dbus.Conn, path string, method string) {
	var res []string

	err := conn.Object("org.ofono", dbus.ObjectPath(path)).Call(method, 0).Store(&res)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to call DBus method:", err)
		return
	}

	fmt.Println("{")

	for _, v := range res {
		fmt.Printf("\t%s\n", v)
	}

	fmt.Println("}")
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
	fmt.Printf("\t%s modem ctx <name> - list modem network contexts\n", args[0])
	fmt.Printf("\t%s modem mon <name> - list servicing cell basic measurements\n", args[0])
	fmt.Printf("\t%s context enable <name> - enable network context\n", args[0])
	fmt.Printf("\t%s context disable <name> - disable network context\n", args[0])
}
