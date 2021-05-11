package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var usersCount map[string]int = map[string]int{
	"shyshlakov":                  50,
	"gore_gorchica":               50,
	"beolnix":                     50,
	"juleswithoutrules":           50,
	"AlenaBetm":                   50,
	"X_tar":                       50,
	"andrey_ya_huy_sredi_blyadey": 50,
	"gdgood":                      50,
	"gleb28":                      50,
	"Ultramarchick":               50,
	"Pesocha":                     50,
	"MaxLyuchin":                  50,
}

func main() {
	rand.Seed(time.Now().UnixNano())
	bot, err := tgbotapi.NewBotAPI("your_token")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	quotes := readfile("./quotes.txt")
	words := readfile("./andrey.txt")

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if update.Message.Text == "/телка" {
			quote := ""
			quoteIndex := random(0, len(quotes)-1)
			quote = quotes[quoteIndex]
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, quote)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}

		if usersCount[update.Message.From.UserName] == 0 {
			usersCount[update.Message.From.UserName] = 50
			wordIndex := random(0, len(words)-1)
			quote := ""
			quote = words[wordIndex]
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, quote)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		} else {
			usersCount[update.Message.From.UserName] = usersCount[update.Message.From.UserName] - 1
		}
	}
}

func readfile(path string) []string {
	buff := make([]string, 0)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buff = append(buff, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return buff
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
