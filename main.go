package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"unicode"

	"github.com/bwmarrin/discordgo"
	"github.com/namsral/flag"
)

var TOKEN string
var USER string
var EMOJI string
var REPEAT bool
var EMOTE bool

func goofy(text string) string {
	var out strings.Builder
	for i, c := range text {
		if i%2 == 0 {
			out.WriteRune(unicode.ToLower(c))
		} else {
			out.WriteRune(unicode.ToUpper(c))
		}
	}
	return out.String()
}

func main() {

	flag.StringVar(&TOKEN, "token", "", "bot token")
	flag.StringVar(&USER, "target", "", "targeted user")
	flag.StringVar(&EMOJI, "emoji", "", "reaction")
	flag.BoolVar(&REPEAT, "repeat", true, "repeat")
	flag.BoolVar(&EMOTE, "emote", true, "emote")
	flag.Parse()

	fmt.Printf("Target: %s, repeat: %t, emote: %t, emoji: %s\n", USER, REPEAT, EMOTE, EMOJI)

	dc, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		panic(err)
	}

	dc.AddHandler(func(s *discordgo.Session, e *discordgo.MessageCreate) {
		if e.Author.ID == USER {
			if EMOTE {
				err := dc.MessageReactionAdd(e.ChannelID, e.ID, EMOJI)
				if err != nil {
					fmt.Println(err)
				}
			}
			if REPEAT {
				_, err = dc.ChannelMessageSend(e.ChannelID, goofy(e.Content))
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	})

	err = dc.Open()
	if err != nil {
		panic(err)
	}
	defer dc.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
