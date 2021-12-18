package main

import (
	"fmt"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  "2115182921:AAEpI_nyJmZi0EGlLKBVlnmJQQfr-3kRnDs",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		if m.Text == "Hello there" {
			b.Send(m.Sender, "General Kenobi")
		}
	})

	b.Handle("/echo", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		b.Send(m.Sender, m.Payload)
		fmt.Println(m.Payload)
	})

	b.Handle("/image", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		switch m.Payload {
		case "screenshoot":
			b.Send(m.Sender, &tb.Photo{File: tb.FromDisk("image.png")})
		}
	})

	fmt.Println("Bot has startted successfully")

	b.Start()

}
