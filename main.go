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
	bot, err := tgbotapi.NewBotAPI("put_your_token_here")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	quotes := readfile()
	words := readAndrfile()

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
			quote = getNextQuote(quotes, len(quotes)-1)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, quote)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}

		if usersCount[update.Message.From.UserName] == 0 {
			usersCount[update.Message.From.UserName] = 50
			rand.Seed(time.Now().UnixNano())
			wordIndex := random(0, len(words)-1)
			quote := ""
			quote = getNextQuote(words, wordIndex)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, quote)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		} else {
			usersCount[update.Message.From.UserName] = usersCount[update.Message.From.UserName] - 1
		}
	}
}

func getNextQuote(quotes []string, quoteIndex int) string {
	if quoteIndex >= len(quotes)-1 {
		quoteIndex = 0
	} else {
		quoteIndex = quoteIndex + 1
	}
	return quotes[quoteIndex]
}

func readfile() []string {
	buff := make([]string, 0)
	file, err := os.Open("./quotes.txt")
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

func readAndrfile() []string {
	buff := make([]string, 0)
	file, err := os.Open("./andrey.txt")
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
