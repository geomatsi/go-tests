package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"time"

	"adsb-bot/adsbtable"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

/* configuration */

type adsbBotConfig struct {
	Addr  string
	Token string
	Chat  int64
}

var botConfig = flag.String("c", "/etc/adsbbot.cfg", "ADS-B Bot configuration file")
var botDebug = flag.Bool("d", false, "ADS-B Bot debug")

/* main */

var adsbLog *adsbtable.AdsbTable
var adsbConf adsbBotConfig

func main() {
	log.SetFlags(0)
	flag.Parse()

	file, err := os.Open(*botConfig)
	defer file.Close()
	if err != nil {
		log.Fatalf("Failed to open config %s: %s\n", *botConfig, err)
	}

	err = json.NewDecoder(file).Decode(&adsbConf)
	if err != nil {
		log.Fatalf("Failed to parse config %s: %s\n", *botConfig, err)
	}

	log.Println("addr: ", adsbConf.Addr)
	log.Println("token: ", adsbConf.Token)
	log.Println("chat: ", adsbConf.Chat)

	adsbLog = adsbtable.NewTable()

	bot := make(chan string)
	go handleBot(bot)

	adsb := make(chan string)
	go handleADSB(adsb)

	cutoff := time.Now()
	beat := time.NewTicker(1 * time.Minute)

EXIT:
	for {
		select {
		case sbs, more := <-adsb:

			if more == false {
				break EXIT
			}

			id, ok := adsbLog.Update(sbs)
			if ok {
				bot <- adsbLog.GetString(id)
			}

		case t := <-beat.C:
			adsbLog.Age(cutoff)
			cutoff = t
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
			adsb, err = net.Dial("tcp", adsbConf.Addr)
			if err != nil {
				log.Printf("ADS-B server dial: %s\n", err)
				time.Sleep(10 * time.Second)
				continue
			} else {
				log.Printf("Connected to ADS-B server %s\n", adsbConf.Addr)
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

			report := tgbotapi.NewMessage(adsbConf.Chat, message)
			_, err := bot.Send(report)
			if err != nil {
				beat = time.NewTimer(5 * time.Second)
				log.Printf("bot failed to send message: %s\n", err)
				conn = false
			}

		case <-beat.C:
			if conn == false {
				bot, err = tgbotapi.NewBotAPI(adsbConf.Token)
				if err != nil {
					beat = time.NewTimer(5 * time.Second)
					log.Printf("Bot connect failed: %s\n", err)
				} else {
					log.Printf("Bot authorized on account %s", bot.Self.UserName)
					bot.Debug = *botDebug
					conn = true
				}
			}
		}
	}
}
