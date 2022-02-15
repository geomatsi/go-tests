package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

/* configuration */

type testBotConfig struct {
	Temp  string
	Token string
	Chat  int64
}

var botConfig = flag.String("c", "/etc/bot.cfg", "Temp Bot configuration file")
var botDebug = flag.Bool("d", false, "Temp Bot debug")

/* main */

var testConf testBotConfig

func main() {
	var upd tgbotapi.UpdatesChannel
	var bot *tgbotapi.BotAPI
	var err error

	log.SetFlags(0)
	flag.Parse()

	file, err := os.Open(*botConfig)
	defer file.Close()
	if err != nil {
		log.Fatalf("Failed to open config %s: %s\n", *botConfig, err)
	}

	err = json.NewDecoder(file).Decode(&testConf)
	if err != nil {
		log.Fatalf("Failed to parse config %s: %s\n", *botConfig, err)
	}

	beat := time.NewTimer(1 * time.Second)
	conn := false

	for {
		select {
		case <-beat.C:
			if conn == false {
				bot, err = tgbotapi.NewBotAPI(testConf.Token)
				if err != nil {
					beat = time.NewTimer(10 * time.Second)
					log.Printf("Bot connect failed: %s\n", err)
					break
				}

				u := tgbotapi.NewUpdate(0)
				upd, err = bot.GetUpdatesChan(u)

				if err != nil {
					beat = time.NewTimer(10 * time.Second)
					log.Printf("Bot GetUpdates failed: %s\n", err)
					break
				}

				if testConf.Chat != 0 {
					report := tgbotapi.NewMessage(testConf.Chat, "Test Bot ready...")
					_, err := bot.Send(report)
					if err != nil {
						beat = time.NewTimer(10 * time.Second)
						log.Printf("bot failed to send message: %s\n", err)
						conn = false
					}
				}

				log.Printf("Bot ready on account %s", bot.Self.UserName)
				bot.Debug = *botDebug
				conn = true
			}

		case update := <-upd:
			if update.Message == nil {
				break
			}

			if update.Message.IsCommand() {
				reply := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				switch update.Message.Command() {
				case "help":
					reply.Text += "Supported commands:\n"
					reply.Text += "\t/status - current test data\n"
				case "toggle":
					fallthrough
				case "list":
					fallthrough
				case "status":
					val, err := readTemp(testConf.Temp)
					if err != nil {
						reply.Text = "Failed to read temperature..."
					} else {
						reply.Text = fmt.Sprintf("%f", float32(val)/1000)
					}
				default:
					reply.Text = "Not yet implemented..."
				}

				_, err := bot.Send(reply)
				if err != nil {
					log.Printf("bot failed to send message: %s\n", err)
					beat = time.NewTimer(10 * time.Second)
					conn = false
				}
			}
		}
	}
}

func readTemp(temp string) (int, error) {
	out, err := exec.Command("cat", temp).Output()
	if err != nil {
		return 0xffff, err
	}

	val, err := strconv.Atoi(string(out[:len(out)-1]))
	if err != nil {
		return 0xffff, err
	}

	return val, nil
}
