package main

import (
	"fmt"
	"os"
	"reflect"

	dbus "github.com/godbus/dbus/v5"
)

func main() {
	conn, err := dbus.SystemBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	var res [][]interface{}

	err = conn.Object("org.ofono", dbus.ObjectPath("/")).Call("org.ofono.Manager.GetModems", 0).Store(&res)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get list of owned names:", err)
		os.Exit(1)
	}

	for _, t := range res {
		switch reflect.TypeOf(t).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(t)

			for i := 0; i < s.Len(); i++ {
				v := s.Index(i)
				fmt.Printf("type[%v] kind[%v] --> %v\n",
					reflect.TypeOf(v), reflect.TypeOf(v).Kind(), v)
			}
		}
	}
}
