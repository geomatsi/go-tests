package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
	"github.com/u-root/u-root/pkg/rtc"
)

var ipAddr = flag.String("a", "192.168.99.57", "SNTP server IP address")
var setTime = flag.Bool("set", false, "Set hwclock from NTP")
var cont = flag.Bool("c", false, "Continuous time readings")
var period = flag.Int("p", 5, "Period in seconds")

func main() {
	log.SetFlags(0)
	flag.Parse()

	rtc0, err := rtc.OpenRTC()
	if err != nil {
		log.Fatalln("Failed to open RTC device: ", err)
	}

	for {
		t1, err := rtc0.Read()
		if err != nil {
			log.Fatalln("Failed to read from RTC device: ", err)
		}

		t2, err := ntp.Time(*ipAddr)
		if err != nil {
			log.Fatalln("Failed to get SNTP time: ", err)
		}

		t3 := time.Now()

		if *setTime {
			err = rtc0.Set(t2)
			if err != nil {
				log.Fatalln("Failed to set RTC from NTP: ", err)
			}

			log.Println("Set RTC from NTP: ", t2.UTC().Format(time.UnixDate))
			os.Exit(0)
		}

		fmt.Printf("\"%s\";\"%s\";\"%s\"\n",
			t1.UTC().Format(time.UnixDate),
			t2.UTC().Format(time.UnixDate),
			t3.UTC().Format(time.UnixDate))

		if !*cont {
			os.Exit(0)
		}

		time.Sleep(time.Duration(*period) * time.Second)
	}
}
