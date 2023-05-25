package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math"
	"strings"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5535936550:AAG2Sv-9C7cE6PHh8QdMc2-giPj8yaIEGso")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			grade := update.Message.Text
			value := 0

			gradesMap := map[uint8]int{
				'A': 5,
				'B': 4,
				'C': 3,
				'D': 2,
				'F': 1,
			}
			//Checks if the grade received is computable, i.e length is 5
			if len(grade) > 7 || len(grade) < 7 {
				reply := tgbotapi.NewMessage(update.Message.Chat.ID, "Wrong format, enter a correct grade")
				bot.Send(reply)
				continue
			} else {
				newGrade := strings.ToUpper(grade)
				N := len(grade)

				for i := 0; i < N; i++ {
					if newGrade[i] != 'F' && gradesMap[newGrade[i]] == 0 {
						reply := tgbotapi.NewMessage(update.Message.Chat.ID, "Wrong format, enter a correct grade")
						bot.Send(reply)
						break
					}
					value += gradesMap[newGrade[i]]
				}
			}

			gradeInNumbers := (float64(value)) / (float64(7))
			num := (gradeInNumbers - 2) / 3
			exp := math.Pow(num, 2.5)

			stipa := (3000) + ((20000 - 3000) * (exp))

			stipa = math.Ceil(stipa)

			//Outputs the grade is is
			response := "Your stipa is: "
			//response += strconv.Itoa(value)
			response += fmt.Sprintf("%v", stipa)
			newMessage := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			bot.Send(newMessage)

			//Bye message
			finalResponse := "Thank you for your usage 😍, "
			finalResponse += update.Message.From.UserName

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, finalResponse)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
