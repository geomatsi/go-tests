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

	// dbus service introspection

	args := os.Args

	fmt.Printf("Introspection: %v -> %v\n", args[1], args[2])
	node, err := introspect.Call(conn.Object(args[1], dbus.ObjectPath(args[2])))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to perform introspection:", err)
		os.Exit(1)
	}

	data, _ := json.MarshalIndent(node, "", "    ")
	os.Stdout.Write(data)
	fmt.Println("")

	// org.ofono.Manager.GetModems: a(oa{sv})

	var modems [][]interface{}

	err = conn.Object("org.ofono", dbus.ObjectPath("/")).Call("org.ofono.Manager.GetModems", 0).Store(&modems)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get the list of available modems:", err)
		os.Exit(1)
	}

	for _, mx := range modems {
		switch reflect.TypeOf(mx).Kind() {
		case reflect.Slice:
			for _, m := range mx {
				switch m.(type) {
				case dbus.ObjectPath:

					// org.ofono.Modem.GetProperties: a{sv}

					var mp map[string]dbus.Variant

					fmt.Printf("Modem %s properties:\n", m)

					err = conn.Object("org.ofono", m.(dbus.ObjectPath)).Call("org.ofono.Modem.GetProperties", 0).Store(&mp)
					if err != nil {
						fmt.Fprintln(os.Stderr, "Failed to get the list of properties:", err)
						os.Exit(1)
					}

					fmt.Println("{")

					for k, v := range mp {
						fmt.Printf("\t%s -> %v\n", k, v)
					}

					fmt.Println("}")

					// org.ofono.cinterion.HardwareMonitor.GetStatistics: a{sv}

					var ms map[string]dbus.Variant

					fmt.Printf("Modem %s stats:\n", m)

					err = conn.Object("org.ofono", m.(dbus.ObjectPath)).Call("org.ofono.cinterion.HardwareMonitor.GetStatistics", 0).Store(&ms)
					if err != nil {
						fmt.Fprintln(os.Stderr, "Failed to get the cinterion stats:", err)
						os.Exit(1)
					}

					fmt.Println("{")

					for k, v := range ms {
						fmt.Printf("\t%s -> %v\n", k, v)
					}

					fmt.Println("}")

					// org.ofono.SimManager.GetProperties: a{sv}

					var sp map[string]dbus.Variant

					fmt.Printf("Modem %s SIM properties:\n", m)

					err = conn.Object("org.ofono", m.(dbus.ObjectPath)).Call("org.ofono.SimManager.GetProperties", 0).Store(&sp)
					if err != nil {
						fmt.Fprintln(os.Stderr, "Failed to get the cinterion stats:", err)
						os.Exit(1)
					}

					fmt.Println("{")

					for k, v := range sp {
						fmt.Printf("\t%s -> %v\n", k, v)
					}

					fmt.Println("}")

				}
			}
		}
	}
}
