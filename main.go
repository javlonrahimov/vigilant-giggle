package main

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

var usersTable string = "userTable"
var wordsTable string = "wordsTable"

func main() {

	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	CreateDatabase()

	CreateTable(usersTable)
	CreateTable(wordsTable)

	b.Handle("/start", func(c tele.Context) error {
		c.Send(c.Sender().Recipient())
		return c.Send("Welcome to the bot.")
	})

	b.Handle("/memo", func(c tele.Context) error {
		InsertWord("id", c.Args()[0], "tableName")
		return c.Send("saved")
	})

	b.Start()

}
