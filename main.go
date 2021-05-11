package main

import (
	"bufio"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("put_your_token_here")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	quotes := readfile()
	words := readAndrfile()
	quoteIndex := 0
	wordIndex := 0

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
			quoteIndex, quote = getNextQuote(quotes, quoteIndex)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, quote)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
		if update.Message.From.UserName == "andrey_ya_huy_sredi_blyadey" ||
			update.Message.From.UserName == "@andrey_ya_huy_sredi_blyadey" ||
			update.Message.From.UserName == "gleb28" ||
			update.Message.From.UserName == "@gleb28" {
			quote := ""
			quoteIndex, quote = getNextQuote(words, wordIndex)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, quote)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}

	}
}

func getNextQuote(quotes []string, quoteIndex int) (int, string) {
	if quoteIndex >= len(quotes)-1 {
		quoteIndex = 0
	} else {
		quoteIndex = quoteIndex + 1
	}

	return quoteIndex, quotes[quoteIndex]
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
