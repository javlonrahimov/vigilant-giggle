package main

import (
	"fmt"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token: "2115182921:AAEpI_nyJmZi0EGlLKBVlnmJQQfr-3kRnDs",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return
	}

	fmt.Println("Bot has startted successfully")

	b.Handle(tb.OnText, func(m *tb.Message) {
		b.Send(m.Sender, "hello world")
	})

	b.Start()
}