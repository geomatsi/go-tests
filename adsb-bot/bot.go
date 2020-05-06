package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/skypies/adsb"
)

var adsbAddr = flag.String("addr", "localhost:30003", "dump1090 SBS-1 service")
var tgDebug = flag.Bool("debug", false, "Telegram Bot debug")
var tgToken = flag.String("token", "00000", "Telegram Bot Token")
var tgChat = flag.Int64("chat", 000000, "Telegram Chat Id")

func main() {
	var msg adsb.Msg

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

			msg.FromSBS1(sbs)

			if msg.HasCallsign() {
				text := msg.Callsign + ": "
				if msg.HasPosition() {
					text += "POS" + msg.Position.String() + " "
				}

				if msg.HasGroundSpeed() {
					text += "GND SPEED(" + strconv.FormatInt(msg.GroundSpeed, 10) + ") "
				}

				if msg.HasVerticalRate() {
					text += "VERT SPEED(" + strconv.FormatInt(msg.VerticalRate, 10) + ") "
				}

				if msg.Altitude != 0 {
					text += "ALT(" + strconv.FormatInt(msg.Altitude, 10) + ") "
				}

				bot <- text
			}
		case t := <-beat.C:
			log.Println("Periodic ADS-B table aging: ", t)
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
