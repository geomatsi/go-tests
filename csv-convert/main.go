package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	csvfile, err := os.Open("test.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	r.Comma = ';'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		rtc, err := time.Parse(time.UnixDate, record[0])
		if err != nil {
			log.Fatalln("failed to parse rtc: ", err)
		}

		ntp, err := time.Parse(time.UnixDate, record[1])
		if err != nil {
			log.Fatalln("failed to parse ntp: ", err)
		}

		sys, err := time.Parse(time.UnixDate, record[2])
		if err != nil {
			log.Fatalln("failed to parse sys: ", err)
		}

		fmt.Printf("\"%s\";\"%d\";\"%d\"\n", ntp, rtc.Unix()-ntp.Unix(), sys.Unix()-ntp.Unix())
	}
}
