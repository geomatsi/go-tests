package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	adsb "github.com/skypies/adsb"
)

/* command line params */

var adsbAddr = flag.String("addr", "localhost:30003", "dump1090 SBS-1 service")
var tgDebug = flag.Bool("debug", false, "Telegram Bot debug")
var tgToken = flag.String("token", "00000", "Telegram Bot Token")
var tgChat = flag.Int64("chat", 000000, "Telegram Chat Id")

/* data */

type adsbTable struct {
	rec map[adsb.IcaoId]adsb.Msg
	syn sync.RWMutex
}

func newTable() *adsbTable {
	tb := &adsbTable{}
	tb.rec = make(map[adsb.IcaoId]adsb.Msg)
	return tb
}

func (tb *adsbTable) update(sbs string) (adsb.IcaoId, bool) {
	var new adsb.Msg
	var cur adsb.Msg

	var tsUpdated bool
	var updated bool
	var ok bool

	tsUpdated = false
	updated = false

	tb.syn.Lock()
	defer tb.syn.Unlock()

	if err := new.FromSBS1(sbs); err != nil {
		return "", false
	}

	id := new.Icao24

	// FIXME: incomplete dump1090 output ?
	if id == "000000" {
		return "", false
	}

	_, ok = tb.rec[id]
	if !ok {
		cur.FromSBS1(fmt.Sprintf(",,,,%s,,,,,,,,,,,,,,,,,", id))
		tb.rec[id] = cur
	}

	cur = tb.rec[id]

	if new.HasCallsign() {
		if cur.Callsign != new.Callsign {
			cur.Callsign = new.Callsign
			updated = true
		}
	}

	if new.HasGroundSpeed() {
		if cur.GroundSpeed != new.GroundSpeed {
			cur.GroundSpeed = new.GroundSpeed
			updated = true
		}
	}

	if new.HasPosition() {
		if cur.Position != new.Position {
			cur.Position = new.Position
			updated = true
		}
	}

	if new.HasVerticalRate() {
		if cur.VerticalRate != new.VerticalRate {
			cur.VerticalRate = new.VerticalRate
			updated = true
		}
	}

	if new.Altitude != 0 {
		if cur.Altitude != new.Altitude {
			cur.Altitude = new.Altitude
			updated = true
		}
	}

	if new.GeneratedTimestampUTC.After(cur.GeneratedTimestampUTC) {
		cur.GeneratedTimestampUTC = new.GeneratedTimestampUTC
		tsUpdated = true
	}

	if updated || tsUpdated {
		tb.rec[id] = cur
	}

	return id, updated
}

func (tb *adsbTable) get(id adsb.IcaoId) (adsb.Msg, bool) {
	tb.syn.Lock()
	defer tb.syn.Unlock()

	ac, ok := tb.rec[id]

	return ac, ok
}

func (tb *adsbTable) getString(id adsb.IcaoId) string {
	var s string

	tb.syn.Lock()
	defer tb.syn.Unlock()

	m, ok := tb.rec[id]
	if !ok {
		return ""
	}

	if m.Icao24 != id {
		fmt.Printf("BOOO: %s != %s\n", m.Icao24, id)
	}

	if m.HasCallsign() {
		s = fmt.Sprintf("%s (%s)", m.Callsign, m.Icao24)
	} else {
		s = fmt.Sprintf("UNKNOWN (%s)", m.Icao24)
	}

	if m.Altitude != 0 {
		s += fmt.Sprintf(" ALT [%d]", m.Altitude)
	}

	if m.HasGroundSpeed() {
		s += fmt.Sprintf(" GND SPEED [%d]", m.GroundSpeed)
	}

	if m.HasVerticalRate() {
		s += fmt.Sprintf(" VERT SPEED [%d]", m.VerticalRate)
	}

	if m.HasPosition() {
		s += fmt.Sprintf(" POS [%s]", m.Position)
	}

	return s
}

func (tb *adsbTable) age(time time.Time) {
	tb.syn.Lock()
	defer tb.syn.Unlock()

	/* TODO */
}

/* main */

var adsbLog *adsbTable

func main() {
	adsbLog = newTable()

	log.SetFlags(0)
	flag.Parse()

	bot := make(chan string)
	go handleBot(bot)

	adsb := make(chan string)
	go handleADSB(adsb)

	beat := time.NewTicker(10 * time.Second)

EXIT:
	for {
		select {
		case sbs, more := <-adsb:

			if more == false {
				break EXIT
			}

			if id, ok := adsbLog.update(sbs); ok {
				bot <- adsbLog.getString(id)
			}

		case <-beat.C:
			// TODO: age ADS-B table
		}
	}
}

func handleADSB(cc chan string) {
	var reader *bufio.Reader
	var adsb net.Conn
	var err error

	conn := false

	for {
		if conn == true {
			str, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					log.Printf("server %s closed connection\n", adsb.RemoteAddr())
				} else {
					log.Printf("adsb server error: %s\n", err)
				}

				conn = false
				continue
			}

			cc <- str
		} else {
			adsb, err = net.Dial("tcp", *adsbAddr)
			if err != nil {
				log.Printf("ADS-B server dial: %s\n", err)
				time.Sleep(10 * time.Second)
				continue
			} else {
				log.Printf("Connected to ADS-B server %s\n", *adsbAddr)
				reader = bufio.NewReader(adsb)
				conn = true
			}
		}
	}
}

func handleBot(cc chan string) {
	var bot *tgbotapi.BotAPI
	var err error

	beat := time.NewTimer(1 * time.Second)
	conn := false

EXIT:
	for {
		select {
		case message, more := <-cc:
			if more == false {
				break EXIT
			}

			log.Printf("New bot message: %s\n", message)

			if conn == false {
				break
			}

			report := tgbotapi.NewMessage(*tgChat, message)
			_, err := bot.Send(report)
			if err != nil {
				beat = time.NewTimer(5 * time.Second)
				log.Printf("bot failed to send message: %s\n", err)
				conn = false
			}

		case <-beat.C:
			if conn == false {
				bot, err = tgbotapi.NewBotAPI(*tgToken)
				if err != nil {
					beat = time.NewTimer(5 * time.Second)
					log.Printf("Bot connect failed: %s\n", err)
				} else {
					log.Printf("Bot authorized on account %s", bot.Self.UserName)
					bot.Debug = *tgDebug
					conn = true
				}
			}
		}
	}
}
