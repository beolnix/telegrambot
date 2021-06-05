package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
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

	goodNight := readfile("./gn.txt")
	goodMorning := readfile("./gm.txt")

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if strings.Contains(update.Message.Text, "/addGNword") {
			result := strings.ReplaceAll(update.Message.Text, "/addGNword", "")
			file, err := os.OpenFile("gn.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatalf("failed creating file: %s", err)
			}
			datawriter := bufio.NewWriter(file)
			_, _ = datawriter.WriteString(result + "\n")
			datawriter.Flush()
			file.Close()
			goodNight = readfile("./gn.txt")
		}
		if strings.Contains(update.Message.Text, "/addGMword") {
			result := strings.ReplaceAll(update.Message.Text, "/addGMword", "")
			file, err := os.OpenFile("gm.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatalf("failed creating file: %s", err)
			}
			datawriter := bufio.NewWriter(file)
			_, _ = datawriter.WriteString(result + "\n")
			datawriter.Flush()
			file.Close()
			goodMorning = readfile("./gm.txt")
		}

		if update.Message.Text == "/телка" {
			quote := ""
			quoteIndex := random(0, len(quotes)-1)
			quote = quotes[quoteIndex]
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, quote)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}

		reg, err := regexp.Compile("[^а-яА-ЯёЁЇїІіЄєҐґ]+")
		if err != nil {
			log.Fatal(err)
		}
		processedString := reg.ReplaceAllString(update.Message.Text, "")
		if contains(goodNight, strings.ToLower(processedString)) {
			indx := random(0, len(goodMorning)-1)
			word := goodMorning[indx]
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, word)
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

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(e, a) {
			return true
		}
	}
	return false
}
